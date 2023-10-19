package entity

// UserDB is an interface for managing user data.
type UserDB interface {
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)

	GetUserRecord(int) (*UserRecordView, error)
	GetAllUsers() ([]*UserItemView, error)
	GetUserById(int) (*UserItemView, error)
	GetAssociationUserByUserId(int) (*UserItemView, error)
	GetCandidateByUserId(int) (*UserItemView, error)
	GetCompanyUserByUserId(int) (*UserItemView, error)
	CreateUser(*UserEntity) (*UserEntity, error)
	CreateAssociationUser(*AssociationUserEntity) (*AssociationUserEntity, error)
	CreateCandidate(*CandidateEntity) (*CandidateEntity, error)
	CreateCompanyUser(*CompanyUserEntity) (*CompanyUserEntity, error)
	AssignUserPhoto(record *UserPhotoEntity) error
	DeleteUserPhoto(userId int) error
	AssignCandidateSkills(candidateId int, records CandidateSkillsEntity) error
	DeleteCandidateSkills(candidateId int) error
	AssignCandidateSpokenLanguages(candidateId int, records CandidateSpokenLanguagesEntity) error
	DeleteCandidateSpokenLanguages(candidateId int) error
	AssignCandidateSeekLocations(candidateId int, records CandidateSeekLocationsEntity) error
	DeleteCandidateSeekLocations(candidateId int) error
	AssignCandidateCV(record *CandidateCVEntity) error
	DeleteCandidateCV(candidateId int) error
	AssignCandidateAttachments(candidateId int, records CandidateAttachmentsEntity) error
	DeleteCandidateAttachments(candidateId int) error
	AssignCandidateVideo(record *CandidateVideoEntity) error
	DeleteCandidateVideo(candidateId int) error
	AssignCandidateEducationHistoryList(candidateId int, records CandidateEducationHistoryListEntity) error
	DeleteCandidateEducationHistoryList(candidateId int) error
	AssignCandidateEmploymentHistoryList(candidateId int, records CandidateEmploymentHistoryListEntity) error
	DeleteCandidateEmploymentHistoryList(candidateId int) error
}
