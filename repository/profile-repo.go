package repository

import (
	"dummy/entities"

	"gorm.io/gorm"
)

type ProfileRepo interface {
	Insert(p entities.Profile) (uint, error)
	Update(p entities.Profile) error
	FindById(ID int64) error
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

func (pr *profileRepo) Update(p entities.Profile) error {
	var err error
	return err
}

func (pr *profileRepo) FindById(profileID int64) error {
	var err error
	return err
}
