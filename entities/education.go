package entities

import "time"

type Education struct {
	ID          int64     `gorm:"primary_key;type:bigint;not null" json:"id"`
	ProfileID   int64     `gorm:"type:bigint; not null" json:"profile_id"`
	School      string    `gorm:"type:varchar(100);not null" json:school"`
	Degree      string    `gorm:"type:varchar(100);not null" json:"degree"`
	StartDate   time.Time `gorm:"type:date;not null;" json:"start_date"`
	EndDate     time.Time `gorm:"type:date;not null;" json:"end_date"`
	City        string    `gorm:"type:varchar(100);not null" json:"city"`
	Description string    `gorm:"type:varchar(100);not null" json:"description"`
	Base
}
