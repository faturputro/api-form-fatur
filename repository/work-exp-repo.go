package repository

import (
	"dummy/entities"

	"gorm.io/gorm"
)

type WorkExperienceRepository interface {
	Add(we entities.WorkExperience, profileID uint) (string, error)
	Get(profileID uint) (string, error)
}

type workExperienceRepository struct {
	db *gorm.DB
}

func NewWorkExperienceRepository(db *gorm.DB) WorkExperienceRepository {
	return &workExperienceRepository{
		db: db,
	}
}

func (wep *workExperienceRepository) Add(we entities.WorkExperience, profileId uint) (string, error) {
	we.ProfileID = profileId
	err := wep.db.Create(&we).Error
	return we.Description, err
}

func (wep *workExperienceRepository) Get(profileId uint) (string, error) {
	var we entities.WorkExperience
	err := wep.db.First(&we, profileId).Error
	return we.Description, err
}
