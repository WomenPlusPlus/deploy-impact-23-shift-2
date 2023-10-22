package entity

var (
	UserKindAdmin       = "ADMIN"
	UserKindAssociation = "ASSOCIATION"
	UserKindCandidate   = "CANDIDATE"
	UserKindCompany     = "COMPANY"
)

var (
	UserRoleAdmin = "ADMIN"
	UserRoleUser  = "USER"
)

var (
	UserStateActive    = "ACTIVE"
	UserStateAnonymous = "ANONYMOUS"
	UserStateDeleted   = "DELETED"
)

var (
	UserJobStatusSearching    = "SEARCHING"
	UserJobStatusOpenTo       = "OPEN_TO"
	UserJobStatusNotSearching = "NOT_SEARCHING"
)

var (
	JobTypeAny        = "ANY"
	JobTypeFullTime   = "FULL_TIME"
	JobTypePartTime   = "PART_TIME"
	JobTypeInternship = "INTERNSHIP"
	JobTypeTemporary  = "TEMPORARY"
)

var (
	CompanySizeAny    = "ANY"
	CompanySizeSmall  = "SMALL"
	CompanySizeMedium = "MEDIUM"
	CompanySizeLarge  = "LARGE"
)

var (
	LocationTypeRemote = "REMOTE"
	LocationTypeHybrid = "HYBRID"
	LocationTypeOnSite = "ON_SITE"
)

var (
	WorkPermitCitizen           = "CITIZEN"
	WorkPermitPermanentResident = "PERMANENT_RESIDENT"
	WorkPermitWorkVisa          = "WORK_VISA"
	WorkPermitStudentVisa       = "STUDENT_VISA"
	WorkPermitTemporary         = "TEMPORARY_RESIDENT"
	WorkPermitNoWorkPermit      = "NO_WORK_PERMIT"
	WorkPermitOther             = "OTHER"
)

var (
	InvitationKindAdmin = "ADMIN"
)
