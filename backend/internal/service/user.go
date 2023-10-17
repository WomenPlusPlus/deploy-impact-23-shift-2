package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"shift/internal/entity"
)

type UserService struct {
	userDB entity.UserDB
}

func NewUserService(userDB entity.UserDB) *UserService {
	return &UserService{
		userDB: userDB,
	}
}

func (s *UserService) CreateCandidate(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	candidate := new(entity.CandidateEntity)
	if err := candidate.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into candidate entity: %w", err)
	}
	logrus.Tracef("Parsed user entity: %+v", candidate.UserEntity)
	logrus.Tracef("Parsed candidate entity: %+v", candidate)

	candidate, err := s.userDB.CreateCandidate(candidate)
	if err != nil {
		return nil, fmt.Errorf("creating new candidate: %w", err)
	}
	logrus.Tracef("Added candidate to db: id=%d", candidate.ID)

	skills := make(entity.CandidateSkillsEntity, 0)
	if err := skills.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into skills entity: %w", err)
	}
	logrus.Tracef("Parsed skills entity: %+v", skills)

	spokenLanguages := make(entity.CandidateSpokenLanguagesEntity, 0)
	if err := spokenLanguages.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into spoken languages entity: %w", err)
	}
	logrus.Tracef("Parsed spoken languages entity: %+v", spokenLanguages)

	seekLocations := make(entity.CandidateSeekLocationsEntity, 0)
	if err := seekLocations.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into seek locations entity: %w", err)
	}
	logrus.Tracef("Parsed seek locations entity: %+v", seekLocations)

	attachments := make(entity.CandidateAttachmentsEntity, 0)
	if err := attachments.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into attachments entity: %w", err)
	}
	logrus.Tracef("Parsed attachments entity: %+v", attachments)

	educationHistoryList := make(entity.CandidateEducationHistoryListEntity, 0)
	if err := educationHistoryList.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into education history list entity: %w", err)
	}
	logrus.Tracef("Parsed education history list entity: %+v", educationHistoryList)

	employmentHistoryList := make(entity.CandidateEmploymentHistoryListEntity, 0)
	if err := employmentHistoryList.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into employment history list entity: %w", err)
	}
	logrus.Tracef("Parsed employment history list entity: %+v", employmentHistoryList)

	if len(skills) > 0 {
		if err := s.userDB.AssignCandidateSkills(candidate.ID, skills); err != nil {
			logrus.Errorf("creating new skills: %v", err)
		} else {
			logrus.Tracef("Added skills to db: id=%d, total=%d", candidate.ID, len(skills))
		}
	} else {
		logrus.Tracef("No skills added to db: id=%d", candidate.ID)
	}

	if len(spokenLanguages) > 0 {
		if err := s.userDB.AssignCandidateSpokenLanguages(candidate.ID, spokenLanguages); err != nil {
			logrus.Errorf("creating new spoken languages: %v", err)
		} else {
			logrus.Tracef("Added spoken languages to db: id=%d, total=%d", candidate.ID, len(spokenLanguages))
		}
	} else {
		logrus.Tracef("No spoken languages added to db: id=%d", candidate.ID)
	}

	if len(seekLocations) > 0 {
		if err := s.userDB.AssignCandidateSeekLocations(candidate.ID, seekLocations); err != nil {
			logrus.Errorf("creating new seek locations: %v", err)
		} else {
			logrus.Tracef("Added seek locations to db: id=%d, total=%d", candidate.ID, len(seekLocations))
		}
	} else {
		logrus.Tracef("No seek locations added to db: id=%d", candidate.ID)
	}

	if len(attachments) > 0 {
		if err := s.userDB.AssignCandidateAttachments(candidate.ID, attachments); err != nil {
			logrus.Errorf("creating new attachments: %v", err)
		} else {
			logrus.Tracef("Added attachments to db: id=%d, total=%d", candidate.ID, len(attachments))
		}
	} else {
		logrus.Tracef("No attachments added to db: id=%d", candidate.ID)
	}

	if len(educationHistoryList) > 0 {
		if err := s.userDB.AssignCandidateEducationHistoryList(candidate.ID, educationHistoryList); err != nil {
			logrus.Errorf("creating new education history list: %v", err)
		} else {
			logrus.Tracef("Added education history list to db: id=%d, total=%d", candidate.ID, len(educationHistoryList))
		}
	} else {
		logrus.Tracef("No education history list added to db: id=%d", candidate.ID)
	}

	if len(employmentHistoryList) > 0 {
		if err := s.userDB.AssignCandidateEmploymentHistoryList(candidate.ID, employmentHistoryList); err != nil {
			logrus.Errorf("creating new employment history list: %v", err)
		} else {
			logrus.Tracef("Added employment history list to db: id=%d, total=%d", candidate.ID, len(employmentHistoryList))
		}
	} else {
		logrus.Tracef("No employment history list added to db: id=%d", candidate.ID)
	}

	return &entity.CreateUserResponse{
		ID:     candidate.ID,
		UserID: candidate.UserID,
	}, nil
}
