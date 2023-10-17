package entity

// UserDB is an interface for managing user data.
type UserDB interface {
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)

	CreateUser(*UserEntity) (*UserEntity, error)
	CreateAssociationUser(*AssociationUserEntity) (*AssociationUserEntity, error)
	CreateCandidate(*CandidateEntity) (*CandidateEntity, error)
	CreateCompanyUser(*CompanyUserEntity) (*CompanyUserEntity, error)
	AssignCandidateSkills(candidateId int, records CandidateSkillsEntity) error
	DeleteCandidateSkills(candidateId int) error
	AssignCandidateSpokenLanguages(candidateId int, records CandidateSpokenLanguagesEntity) error
	DeleteCandidateSpokenLanguages(candidateId int) error
	AssignCandidateSeekLocations(candidateId int, records CandidateSeekLocationsEntity) error
	DeleteCandidateSeekLocations(candidateId int) error
	AssignCandidateAttachments(candidateId int, records CandidateAttachmentsEntity) error
	DeleteCandidateAttachments(candidateId int) error
	AssignCandidateEducationHistoryList(candidateId int, records CandidateEducationHistoryListEntity) error
	DeleteCandidateEducationHistoryList(candidateId int) error
	AssignCandidateEmploymentHistoryList(candidateId int, records CandidateEmploymentHistoryListEntity) error
	DeleteCandidateEmploymentHistoryList(candidateId int) error
}
