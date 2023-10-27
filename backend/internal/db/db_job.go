package db

import (
	"shift/internal/entity"

	_ "github.com/lib/pq"
)

// CreateJob implements entity.JobDB.
func (*PostgresDB) CreateJob(*entity.JobEntity) (*entity.JobEntity, error) {
	panic("unimplemented")
}

// DeleteJob implements entity.JobDB.
func (*PostgresDB) DeleteJob(int) error {
	panic("unimplemented")
}

// GetAllJobs implements entity.JobDB.
func (*PostgresDB) GetAllJobs() ([]*entity.JobItemView, error) {
	panic("unimplemented")
}

// GetJobById implements entity.JobDB.
func (*PostgresDB) GetJobById(int) (*entity.JobItemView, error) {
	panic("unimplemented")
}
