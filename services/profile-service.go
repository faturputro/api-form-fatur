package services

import (
	"dummy/cache"
	"dummy/entities"
	"dummy/repository"
	"errors"
)

type ProfileService interface {
	Create(p entities.Profile) (uint, error)
	Update(p entities.Profile, id uint) (uint, error)
	Get(profileID uint) (entities.Profile, error)
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

func (ps *profileService) Update(p entities.Profile, id uint) (uint, error) {
	var err error
	found, err := ps.repo.FindById(id)
	if found.ID == 0 {
		err = errors.New("Profile not found")
		return 0, err
	}
	err = ps.repo.Update(p, id)
	return id, err
}

func (ps *profileService) Get(profileId uint) (entities.Profile, error) {
	profile, err := ps.repo.FindById(profileId)
	return profile, err
}
