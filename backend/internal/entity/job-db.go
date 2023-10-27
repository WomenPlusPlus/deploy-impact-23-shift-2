package entity

// JobDB is an interface for managing  jobs data.
type JobDB interface {
	CreateJob(*JobEntity) (*JobEntity, error)
	DeleteJob(int) error
	GetAllJobs() ([]*JobItemView, error)
	GetJobById(int) (*JobItemView, error)
}
