package entity

import (
	"time"
)

type JobListing struct {
	ID                     int    `json:"id"`
	CompanyId              int    `json:"companyId"`
	Title                  string `json:"title"`
	Experience             string `json:"experience"`
	JobType                string `json:"jobType"`         // from predef
	EmploymentLevel        string `json:"employmentLevel"` // 80 % 100% ...
	Overview               string `json:"overview"`        //overview
	RoleAndResponibilities string `json:"roleResponibilities"`
	NiceToHave             string `json:"niceToHave"`
	CandidateDescription   string `json:"candidateDescription"` //Who you are from interface
	//SkillsRequired         string    `json:"skills"`    // how to make skills and experience list ? skill1 experience level ?
	// most important 3 skills and expeience level - to use for matching (?)
	Skills []JobSkill `json:"skills"`
	// Skill1           string    `json:"skill1"`           // should be from predefined set ?
	// ExperienceSkill1 string    `json:experienceSkill1"`  // should be from predefined set
	// Skill2           string    `json:"skill2"`           // should be from predefined set?
	// ExperienceSkill2 string    `json:"experienceSkill2"` // should be from predefined set
	// Skill3           string    `json:"skill3"`           // should be from predefined set?
	// ExperienceSkill3 string    `json:"experienceSkill3"` // should be from predefined set
	SpokenLanguages []SpokenLanguage `json:"spokenLanguages"` // spoken languages, should be from predefined set
	Location        string           `json:"location"`
	LocationType    string           `sjon:"locationType"` //  hybrid remode, .. from predef
	SalaryRange     string           `json:"salaryRange"`
	Benefits        string           `json:"benefits"`
	StartDate       time.Time        `json:"startDate"`       // type?
	CreatedByUserId int              `json:"createdByUserId"` // user ref company user
	Status          bool             `json:"status"`
	CreatedAt       time.Time        `json:"createdAt"`
}

// to change
type JobListingEntity struct {
	ID                     int    `json:"id"`
	CompanyId              int    `json:"company_id"`
	Title                  string `json:"title"`
	Experience             string `json:"experience"`
	JobType                string `json:"job_type"`         // from predef
	EmploymentLevel        string `json:"employment_level"` // 80 % 100% ...
	Overview               string `json:"overview"`         //overview
	RoleAndResponibilities string `json:"role_responibilities"`
	NiceToHave             string `json:"nice_ho_have"`
	CandidateDescription   string `json:"candidate_description"` //Who you are from interface
	//SkillsRequired         string    `json:"skills"`    // how to make skills and experience list ? skill1 experience level ?
	// most important 3 skills and expeience level - to use for matching (?)
	Skills []JobSkill `json:"skills"`
	// Skill1           string    `json:"skill1"`           // should be from predefined set ?
	// ExperienceSkill1 string    `json:experienceSkill1"`  // should be from predefined set
	// Skill2           string    `json:"skill2"`           // should be from predefined set?
	// ExperienceSkill2 string    `json:"experienceSkill2"` // should be from predefined set
	// Skill3           string    `json:"skill3"`           // should be from predefined set?
	// ExperienceSkill3 string    `json:"experienceSkill3"` // should be from predefined set
	SpokenLanguages []SpokenLanguage `json:"spoken_languages"` // spoken languages, should be from predefined set
	Location        string           `json:"location"`
	LocationType    string           `sjon:"location_type"` //  hybrid remode, .. from predef
	SalaryRange     string           `json:"salary_range"`
	Benefits        string           `json:"benefits"`
	StartDate       time.Time        `json:"start_date"`         // type?
	CreatedByUserId int              `json:"created_by_user_id"` // user ref company user
	Status          bool             `json:"status"`
	CreatedAt       time.Time        `json:"created_at"`
}

// todo
func NewJobListing(title, experience, workType, employmentLevel, overview, roleAndResponibilities, niceToHave,
	candidateDescription, skill1, experienceSkill1, skill2, experienceSkill2, skill3, experienceSkill3,
	languages, location, salaryRange, benefits string,
	companyId, createdByUserId int,
	status bool,
	createdAt, startDate time.Time) *JobListing {
	return &JobListing{
		ID:                     companyId,
		CompanyId:              companyId,
		Title:                  title,
		Experience:             experience,
		JobType:                workType,
		EmploymentLevel:        employmentLevel,
		Overview:               overview,
		RoleAndResponibilities: roleAndResponibilities,
		NiceToHave:             niceToHave,
		CandidateDescription:   candidateDescription,
		Skills:                 []JobSkill{},
		SpokenLanguages:        []SpokenLanguage{},
		Location:               location,
		LocationType:           location,
		SalaryRange:            salaryRange,
		Benefits:               benefits,
		StartDate:              startDate,
		CreatedByUserId:        createdByUserId,
		Status:                 status,
		CreatedAt:              createdAt,
	}

}

// // tmp:
// func NewJobListing(title, overview string) *JobListing {
// 	return &JobListing{
// 		CompanyId:              1,                                               //                int    `json:"companyId"`
// 		Title:                  title,                                           //            string `json:"title"`
// 		Experience:             "x",                                             //       string `json:"experience"`
// 		WorkType:               "x",                                             //  string `json:"workType"`        // internship etc
// 		EmploymentLevel:        "x",                                             //        string `json:"employmentLevel"` // 80 % 100% ...
// 		Overview:               overview,                                        //          string `json:"overview"`        //overview
// 		RoleAndResponibilities: "x",                                             //  string `json:"roleAndResponibilities"`
// 		NiceToHave:             "x",                                             // string `json:"niceToHave"`
// 		CandidateDescription:   "x",                                             //   string `json:"candidateDescription"` //Who you are from interface
// 		Skill1:                 "x",                                             //        string    `json:"skill1"`           // should be from predefined set ?
// 		ExperienceSkill1:       "x",                                             // string    `json:eExperienceSkill1"` // should be from predefined set
// 		Skill2:                 "x",                                             //        string    `json:"skill2"`           // should be from predefined set?
// 		ExperienceSkill2:       "x",                                             // string    `json:"experienceSkill2"` // should be from predefined set
// 		Skill3:                 "x",                                             //      string    `json:"skill3"`           // should be from predefined set?
// 		ExperienceSkill3:       "x",                                             //  string    `json:"experienceSkill3"` // should be from predefined set
// 		Languages:              "x",                                             //       string    `json:"languages"`        // spoken languages, should be from predefined set
// 		Location:               "x",                                             //        string    `json:"location"`
// 		SalaryRange:            "x",                                             //     string    `json:"salaryRange"`
// 		Benefits:               "x",                                             //        string    `json:"benefits"`
// 		StartDate:              time.Date(2023, 11, 12, 0, 0, 0, 0, time.Local), //        time.Time `json:"startDate"`       // type?
// 		CreatedByUserId:        1,                                               //  int       `json:"createdByUserId"` // user ref company user
// 		Status:                 1,
// 		CreatedAt:              time.Now().UTC(),
// 	}
// }

func (u *JobListingEntity) FromCreationRequest(request *CreateUserRequest) error {
	//todo
	return nil
}
