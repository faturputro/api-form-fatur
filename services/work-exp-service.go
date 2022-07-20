package services

import (
	"dummy/entities"
	"dummy/repository"
)

type WorkExperienceService interface {
	Add(we entities.WorkExperience, profileId uint) (string, error)
	Get(profileId uint) (string, error)
}

type workingExperienceService struct {
	repo repository.WorkExperienceRepository
}

func NewWorkExperienceService(repo repository.WorkExperienceRepository) WorkExperienceService {
	return &workingExperienceService{
		repo: repo,
	}
}

func (wes *workingExperienceService) Add(we entities.WorkExperience, profileId uint) (string, error) {
	return wes.repo.Add(we, profileId)
}

func (wes *workingExperienceService) Get(profileId uint) (string, error) {
	return wes.repo.Get(profileId)
}
