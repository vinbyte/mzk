package student

import (
	"context"
	"time"

	"github.com/vinbyte/mzk/configs"
)

type StudentService interface {
	GetRecords(ctx context.Context, filter RecordFilter) (records []Record, err error)
}

// StudentServiceImpl is the service implementation for Student entities.
type StudentServiceImpl struct {
	StudentRepository StudentRepository
	Config            *configs.Config
}

// ProvideStudentServiceImpl is the provider for this service.
func ProvideStudentServiceImpl(studentRepository StudentRepository, config *configs.Config) *StudentServiceImpl {
	s := new(StudentServiceImpl)
	s.StudentRepository = studentRepository
	s.Config = config
	return s
}

func (s *StudentServiceImpl) GetRecords(ctx context.Context, filter RecordFilter) (records []Record, err error) {
	records, err = s.StudentRepository.FetchByFilter(ctx, filter)
	if err != nil {
		return
	}

	// format created at to UTC RFC 3339
	for i := 0; i < len(records); i++ {
		records[i].CreatedAtFormatted = records[i].CreatedAt.In(time.UTC).Format(time.RFC3339)
	}

	return
}
