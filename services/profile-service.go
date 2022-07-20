package services

import (
	"dummy/cache"
	"dummy/entities"
	"dummy/repository"
)

type ProfileService interface {
	Create(p entities.Profile) (uint, error)
	Update(p entities.Profile) (interface{}, error)
	Get(profileID int64) (entities.Profile, error)
}

type profileService struct {
	cache cache.AppCache
	repo  repository.ProfileRepo
}

func NewProfileService(c cache.AppCache, r repository.ProfileRepo) ProfileService {
	return &profileService{
		cache: c,
		repo:  r,
	}
}

func (ps *profileService) Create(p entities.Profile) (uint, error) {
	var err error
	profileId, err := ps.repo.Insert(p)
	return profileId, err
}

func (ps *profileService) Update(p entities.Profile) (interface{}, error) {
	var err error
	return nil, err
}

func (ps *profileService) Get(profileId int64) (entities.Profile, error) {
	var (
		err     error
		profile entities.Profile
	)
	return profile, err
}
