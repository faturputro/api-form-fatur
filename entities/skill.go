package entities

type Skill struct {
	ID        int64  `gorm:"primary_key;type:bigint;not null" json:"id"`
	ProfileID int64  `gorm:"type:bigint; not null" json:"profile_id"`
	Skill     string `gorm:"type:varchar(100);not null" json:skill"`
	Level     string `gorm:"type:varchar(100);not null" json:"level"`
	Base
}
