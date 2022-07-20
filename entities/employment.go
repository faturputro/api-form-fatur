package entities

import "time"

type Employment struct {
	ID          int64     `gorm:"primary_key;type:bigint;not null" json:"id"`
	ProfileID   int64     `gorm:"type:bigint; not null" json:"profile_id"`
	JobTitle    string    `gorm:"type:varchar(100);not null" json:job_title"`
	Employer    string    `gorm:"type:varchar(100);not null" json:"employer"`
	StartDate   time.Time `gorm:"type:date;not null;" json:"start_date"`
	EndDate     time.Time `gorm:"type:date;not null;" json:"end_date"`
	City        string    `gorm:"type:varchar(100);not null" json:"city"`
	Description string    `gorm:"type:varchar(100);not null" json:"description"`
	Base
}
