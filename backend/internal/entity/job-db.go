package entity

type JobDB interface {
	CreateJob(job *JobEntity) (*JobEntity, error)
	AssignJobLocation(jobId int, location *JobLocationEntity) error
	AssignJobSkills(jobId int, skills JobSkillsEntity) error
	AssignJobLanguages(jobId int, locations JobLanguagesEntity) error
	DeleteJob(int) error
	GetAllJobs() ([]*JobView, error)
	GetJobById(int) (*JobView, error)
	GetSkillsByJobId(int) (JobSkillsEntity, error)
	GetLanguagesByJobId(int) (JobLanguagesEntity, error)
}
