package service

import (
	"context"
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
)

type JobService struct {
	bucketDB       entity.BucketDB
	jobDB          entity.JobDB
	userService    *UserService
	companyService *CompanyService
}

func NewJobService(
	bucketDB entity.BucketDB,
	jobDB entity.JobDB,
	userService *UserService,
	companyService *CompanyService,
) *JobService {
	return &JobService{
		bucketDB:       bucketDB,
		jobDB:          jobDB,
		userService:    userService,
		companyService: companyService,
	}
}

func (s *JobService) CreateJob(ctx context.Context, req *entity.CreateJobRequest) (*entity.CreateJobResponse, error) {
	job := new(entity.JobEntity)
	if err := job.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into job entity: %w", err)
	}
	if creatorId, ok := ctx.Value(entity.ContextKeyUserId).(int); ok {
		job.CreatorID = creatorId
	} else {
		return nil, fmt.Errorf("parsing creator id for job entity: got=%v", ctx.Value(entity.ContextKeyUserId))
	}
	logrus.Tracef("Parsed job entity: %+v", job)

	job, err := s.jobDB.CreateJob(job)
	if err != nil {
		return nil, fmt.Errorf("creating new job: %w", err)
	}
	logrus.Tracef("Added job to db: id=%d", job.ID)

	location := new(entity.JobLocationEntity)
	if err := location.FromCreationRequest(job.ID, req); err != nil {
		return nil, fmt.Errorf("parsing request into job location entity: %w", err)
	}
	logrus.Tracef("Parsed job location entity: %+v", location)

	skills := make(entity.JobSkillsEntity, 0)
	if err := skills.FromCreationRequest(job.ID, req); err != nil {
		return nil, fmt.Errorf("parsing job request into skills entity: %w", err)
	}
	logrus.Tracef("Parsed job skills entity: %+v", skills)

	languages := make(entity.JobLanguagesEntity, 0)
	if err := languages.FromCreationRequest(job.ID, req); err != nil {
		return nil, fmt.Errorf("parsing request into job languages entity: %w", err)
	}
	logrus.Tracef("Parsed job languages entity: %+v", languages)

	if err := s.jobDB.AssignJobLocation(job.ID, location); err != nil {
		logrus.Errorf("creating job location: %v", err)
	} else {
		logrus.Tracef("Added job location to db: id=%d", job.ID)
	}

	if len(skills) > 0 {
		if err := s.jobDB.AssignJobSkills(job.ID, skills); err != nil {
			logrus.Errorf("creating new skills: %v", err)
		} else {
			logrus.Tracef("Added job skills to db: id=%d, total=%d", job.ID, len(skills))
		}
	} else {
		logrus.Tracef("No job skills added to db: id=%d", job.ID)
	}

	if len(languages) > 0 {
		if err := s.jobDB.AssignJobLanguages(job.ID, languages); err != nil {
			logrus.Errorf("creating new job languages: %v", err)
		} else {
			logrus.Tracef("Added job languages to db: id=%d, total=%d", job.ID, len(languages))
		}
	} else {
		logrus.Tracef("No job languages added to db: id=%d", job.ID)
	}

	return &entity.CreateJobResponse{ID: job.ID}, nil
}

func (s *JobService) GetAllJobs() (*entity.ListJobsResponse, error) {
	jobs, err := s.jobDB.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("getting all jobs: %w", err)
	}
	logrus.Tracef("Get all jobs from db: total=%d", len(jobs))

	res := new(entity.ListJobsResponse)
	res.Items = make([]*entity.JobItemResponse, len(jobs))
	for i, job := range jobs {
		item, err := s.getJobItemResponse(job)
		if err != nil {
			return nil, err
		}
		res.Items[i] = item
	}
	return res, nil
}

func (s *JobService) GetJobById(id int) (*entity.ViewJobResponse, error) {
	job, err := s.jobDB.GetJobById(id)
	if err != nil {
		return nil, fmt.Errorf("getting all jobs: %w", err)
	}
	logrus.Tracef("Get job from db: id=%d", job.ID)

	return s.getViewJobResponse(job)
}

func (s *JobService) getJobItemResponse(job *entity.JobView) (*entity.JobItemResponse, error) {
	res := new(entity.JobItemResponse)
	res.FromJobView(job)

	user, err := s.userService.GetUserByCompanyUserId(job.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("could not get company by company user %d for job %d: %w", job.CreatorID, job.ID, err)
	}
	res.Creator = new(entity.ViewJobCreator)
	res.Creator.FromViewUserResponse(user)

	company, err := s.companyService.GetCompanyByCompanyUserId(job.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("could not get company by company user %d for job %d: %w", job.CreatorID, job.ID, err)
	}
	res.Company = new(entity.ViewJobCompany)
	res.Company.FromViewCompanyResponse(company)
	return res, nil
}

func (s *JobService) getViewJobResponse(job *entity.JobView) (*entity.ViewJobResponse, error) {
	res := new(entity.ViewJobResponse)
	res.FromJobView(job)

	skills, err := s.jobDB.GetSkillsByJobId(job.ID)
	if err != nil {
		return nil, fmt.Errorf("could not get skills for job %d: %w", job.ID, err)
	}
	res.FromSkillsEntity(skills)

	languages, err := s.jobDB.GetLanguagesByJobId(job.ID)
	if err != nil {
		return nil, fmt.Errorf("could not get languages for job %d: %w", job.ID, err)
	}
	res.FromLanguagesEntity(languages)

	user, err := s.userService.GetUserByCompanyUserId(job.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("could not get company by company user %d for job %d: %w", job.CreatorID, job.ID, err)
	}
	res.Creator = new(entity.ViewJobCreator)
	res.Creator.FromViewUserResponse(user)

	company, err := s.companyService.GetCompanyByCompanyUserId(job.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("could not get company by company user %d for job %d: %w", job.CreatorID, job.ID, err)
	}
	res.Company = new(entity.ViewJobCompany)
	res.Company.FromViewCompanyResponse(company)
	return res, nil
}
