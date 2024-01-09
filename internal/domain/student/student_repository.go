package student

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/vinbyte/mzk/infras"
	"github.com/vinbyte/mzk/shared/failure"
	"github.com/vinbyte/mzk/shared/logger"
)

var (
	studentQueries = struct {
		selectRecords string
	}{
		selectRecords: `
		WITH r AS (
			SELECT 
				id,
				created_at,
				SUM(m) AS total_marks
			FROM records, unnest(marks) AS m
			GROUP BY id
		)
		SELECT 
			r.id,
			r.created_at,
			r.total_marks
		FROM r
		WHERE true
		`,
	}
)

// StudentRepository is the repository for student data.
type StudentRepository interface {
	FetchByFilter(ctx context.Context, filter RecordFilter) (records []Record, err error)
}

// StudentRepositoryPostgres is the MySQL-backed implementation of FooRepository.
type StudentRepositoryPostgres struct {
	DB *infras.PostgresConn
}

// ProvideStudentRepositoryPostgres is the provider for this repository.
func ProvideStudentRepositoryPostgres(db *infras.PostgresConn) *StudentRepositoryPostgres {
	s := new(StudentRepositoryPostgres)
	s.DB = db

	return s
}

// FetchByFilter resolves a Records by filter
func (s *StudentRepositoryPostgres) FetchByFilter(ctx context.Context, filter RecordFilter) (records []Record, err error) {
	whereClause, args := s.buildWhereClause(filter)
	query := studentQueries.selectRecords + whereClause
	err = s.DB.Read.SelectContext(
		ctx,
		&records,
		query,
		args...)
	if err != nil {
		logger.ErrorWithStack(err)
		if err == sql.ErrNoRows {
			err = nil
		} else {
			err = failure.InternalError(errors.New(ErrInternalServerErrorMsg))
		}
	}

	return
}

func (s *StudentRepositoryPostgres) buildWhereClause(filter RecordFilter) (whereClause string, args []interface{}) {
	counter := 0

	// if start time and end time is available, add it to where clause
	if !filter.StartTime.IsZero() && !filter.EndTime.IsZero() {
		// build query like this form : and (r.created_at between '2024-01-09 00:00:00'::timestamptz and '2024-01-09 23:59:59'::timestamptz)
		startTimeFormatted := filter.StartTime.Format("2006-01-02 15:04:05")
		endTimeFormatted := filter.EndTime.Format("2006-01-02 15:04:05")
		counter++
		startTimeIndexParam := counter
		counter++
		endTimeIndexParam := counter
		whereClause += fmt.Sprintf(" AND (r.created_at between $%d::timestamptz AND $%d::timestamptz)", startTimeIndexParam, endTimeIndexParam)
		args = append(args, startTimeFormatted, endTimeFormatted)
	}

	// if filter min count and max count is available, add it to where clause
	if filter.MinCount > 0 && filter.MaxCount > 0 {
		counter++
		minCountIndexParam := counter
		counter++
		maxCountIndexParam := counter
		whereClause += fmt.Sprintf(" AND (r.total_marks between $%d and $%d)", minCountIndexParam, maxCountIndexParam)
		args = append(args, filter.MinCount, filter.MaxCount)
	}

	return
}
