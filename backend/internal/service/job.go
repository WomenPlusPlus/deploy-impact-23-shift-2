package service

import (
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
)

type JobService struct {
	bucketDB entity.BucketDB
	jobDB    entity.JobDB
}

func NewJobService(bucketDB entity.BucketDB, jobDB entity.JobDB) *JobService {
	return &JobService{
		bucketDB: bucketDB,
		jobDB:    jobDB,
	}
}

func (s *JobService) CreateJob(req *entity.CreateJobRequest) (*entity.CreateJobResponse, error) {

	//s.createJob(req)

}

func (s *JobService) ListJobs() (*entity.ListJobsResponse, error) {
	jobs, err := s.jobDB.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("getting all jobs: %w", err)
	}
	logrus.Tracef("Get all jobs from db: total=%d", len(jobs))

	res := new(entity.ListJobsResponse)
	res.FromJobsView(jobs)
	return res, nil
}
