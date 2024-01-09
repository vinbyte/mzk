package student

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	FilterTimeFormat = "2006-01-02"

	ErrInternalServerErrorMsg              = "internal server error"
	ErrMinCountCannotNegative              = "minCount cannot negative"
	ErrMaxCountCannotNegative              = "maxCount cannot negative"
	ErrStartDateMustLessOrEqualThanEndDate = "startDate must less or equal than endDate"
)

// Record is a representation of records table
type Record struct {
	ID                 int64     `db:"id" json:"id"`
	Name               string    `db:"name" json:"-"`
	Marks              []int     `db:"marks" json:"-"`
	CreatedAt          time.Time `db:"created_at" json:"-"`
	CreatedAtFormatted string    `db:"-" json:"createdAt"`
	TotalMarks         int       `db:"total_marks" json:"totalMarks"`
}

// RecordFilter is a filter data for records table
type RecordFilter struct {
	StartDate string    `json:"startDate" example:"2024-01-09"`
	StartTime time.Time `json:"-"`
	EndDate   string    `json:"endDate" example:"2024-01-10"`
	EndTime   time.Time `json:"-"`
	MinCount  int       `json:"minCount" example:"100"`
	MaxCount  int       `json:"maxCount" example:"300"`
}

func (r *RecordFilter) Validate() (err error) {
	if r.StartDate != "" {
		r.StartTime, err = time.Parse(FilterTimeFormat, r.StartDate)
		if err != nil {
			log.Warn().Str("startDate", r.StartDate).Err(err)
			err = fmt.Errorf("startDate error : %s", err.Error())

			return
		}
		r.StartTime = r.StartTime.UTC()
	}
	if r.EndDate != "" {
		r.EndTime, err = time.Parse(FilterTimeFormat, r.EndDate)
		if err != nil {
			log.Warn().Str("endDate", r.EndDate).Err(err)
			err = fmt.Errorf("endDate error : %s", err.Error())

			return
		}
		// set time to 23:59:59.999
		r.EndTime = r.EndTime.UTC().AddDate(0, 0, 1).Add(-1 * time.Millisecond)
	}
	// if time range available, make sure start time less than end time
	if !r.StartTime.IsZero() && !r.EndTime.IsZero() && r.StartTime.After(r.EndTime) {
		err = errors.New(ErrStartDateMustLessOrEqualThanEndDate)
		return
	}

	if r.MinCount < 0 {
		err = errors.New(ErrMinCountCannotNegative)
		return
	}
	if r.MaxCount < 0 {
		err = errors.New(ErrMaxCountCannotNegative)
	}

	return
}

type RecordData struct {
	Records []Record `json:"records"`
}
