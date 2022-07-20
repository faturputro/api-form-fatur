package repository

import (
	"dummy/entities"

	"gorm.io/gorm"
)

type ProfileRepo interface {
	Insert(p entities.Profile) (uint, error)
	Update(p entities.Profile, id uint) error
	FindById(ID uint) (entities.Profile, error)
}

type profileRepo struct {
	db *gorm.DB
}

func NewProfileRepo(db *gorm.DB) ProfileRepo {
	return &profileRepo{
		db: db,
	}
}

func (pr *profileRepo) Insert(p entities.Profile) (uint, error) {
	var err error
	err = pr.db.Create(&p).Error
	return p.ID, err
}

func (pr *profileRepo) Update(p entities.Profile, id uint) error {
	p.ID = id
	return pr.db.Save(&p).Error
}

func (pr *profileRepo) FindById(profileID uint) (entities.Profile, error) {
	var (
		err     error
		profile entities.Profile
	)
	err = pr.db.First(&profile, profileID).Error
	return profile, err
}
