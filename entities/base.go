package entities

import (
	"time"
)

type Base struct {
	Created        time.Time   `gorm:"type:timestamp;not null;default:timezone('utc', now())" json:"created"`
	CreatedBy      interface{} `gorm:"type:bigint;" json:"created_by"`
	LastModified   interface{} `gorm:"type:timestamp;" json:"last_modified"`
	LastModifiedBy interface{} `gorm:"type:bigint;" json:"last_modified_by"`
}
