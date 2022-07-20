package entities

import "gorm.io/gorm"

type WorkExperience struct {
	gorm.Model
	Description string `gorm:"type:text;not null" json:"description"`
	ProfileID   uint   `gorm:"type:int; not null" json:"profile_id"`
	Base
}
