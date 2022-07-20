package entities

type WorkExperience struct {
	ID          int64  `gorm:"primary_key;type:bigint;not null" json:"id"`
	Description string `gorm:"type:text;not null" json:"description"`
	ProfileID   int64  `gorm:"type:bigint; not null" json:"profile_id"`
	Base
}
