package entities

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	// ID               int64     `gorm:"primary_key;type:bigint;AUTO_INCREMENT" json:"id"`
	FirstName        string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName         string    `gorm:"type:varchar(255)" json:"last_name"`
	Email            string    `gorm:"type:varchar(255);unique; not null" json:"email"`
	Phone            string    `gorm:"type:varchar(15);unique;not null" json:"phone"`
	PlaceOfBirth     string    `gorm:"type:varchar(50)" json:"place_of_birth"`
	DateOfBirth      time.Time `gorm:"type:date" json:"date_of_birth"`
	WantedJobTitle   string    `gorm:"type:varchar(50)" json:"wanted_job_title"`
	City             string    `gorm:"type:varchar(50)" json:"city"`
	PostalCode       int64     `gorm:"type:bigint" json:"postal_code"`
	DrivingLicenseNo string    `gorm:"type:varchar(50)" json:"driving_license_no"`
	Nationality      string    `gorm:"type:varchar(50)" json:"nationality"`
	Country          string    `gorm:"type:varchar(50)" json:"country"`
	Address          string    `gorm:"type:text" json:"address"`
	Base
}
