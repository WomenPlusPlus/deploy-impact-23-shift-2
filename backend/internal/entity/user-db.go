package entity

// UserDB is an interface for managing user data.
type UserDB interface {
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)

<<<<<<< .merge_file_F7KVQJ
	GetUserRecord(int) (*UserRecordView, error)
	GetAllUsers() ([]*UserItemView, error)
	GetUserById(int) (*UserItemView, error)
	GetAssociationUserByUserId(int) (*UserItemView, error)
	GetCandidateByUserId(int) (*UserItemView, error)
	GetCompanyUserByUserId(int) (*UserItemView, error)
=======
	GetAllUsers() ([]*UserItemView, error)
>>>>>>> .merge_file_pRpMbJ
	CreateUser(*UserEntity) (*UserEntity, error)
	EditUser(int, *UserEntity) (*UserEntity, error)
	CreateAssociationUser(*AssociationUserEntity) (*AssociationUserEntity, error)
	EditAssociationUser(int, *AssociationUserEntity) (*AssociationUserEntity, error)
	CreateCandidate(*CandidateEntity) (*CandidateEntity, error)
	EditCandidate(int, *CandidateEntity) (*CandidateEntity, error)
	CreateCompanyUser(*CompanyUserEntity) (*CompanyUserEntity, error)
	EditCompanyUser(int, *CompanyUserEntity) (*CompanyUserEntity, error)
	AssignUserPhoto(record *UserPhotoEntity) error
	DeleteUserPhoto(userId int) error
	GetCandidateSkills(candidateId int) (CandidateSkillsEntity, error)
	AssignCandidateSkills(candidateId int, records CandidateSkillsEntity) error
	DeleteCandidateSkills(candidateId int) error
	GetCandidateSpokenLanguages(candidateId int) (CandidateSpokenLanguagesEntity, error)
	AssignCandidateSpokenLanguages(candidateId int, records CandidateSpokenLanguagesEntity) error
	DeleteCandidateSpokenLanguages(candidateId int) error
	GetCandidateSeekLocations(candidateId int) (CandidateSeekLocationsEntity, error)
	AssignCandidateSeekLocations(candidateId int, records CandidateSeekLocationsEntity) error
	DeleteCandidateSeekLocations(candidateId int) error
	AssignCandidateCV(record *CandidateCVEntity) error
	DeleteCandidateCV(candidateId int) error
	GetCandidateAttachments(candidateId int) (CandidateAttachmentsEntity, error)
	AssignCandidateAttachments(candidateId int, records CandidateAttachmentsEntity) error
	DeleteCandidateAttachments(candidateId int) error
	AssignCandidateVideo(record *CandidateVideoEntity) error
	DeleteCandidateVideo(candidateId int) error
	GetCandidateEducationHistoryList(candidateId int) (CandidateEducationHistoryListEntity, error)
	AssignCandidateEducationHistoryList(candidateId int, records CandidateEducationHistoryListEntity) error
	DeleteCandidateEducationHistoryList(candidateId int) error
	GetCandidateEmploymentHistoryList(candidateId int) (CandidateEmploymentHistoryListEntity, error)
	AssignCandidateEmploymentHistoryList(candidateId int, records CandidateEmploymentHistoryListEntity) error
	DeleteCandidateEmploymentHistoryList(candidateId int) error
}
