package entity

import (
	"time"

	"gorm.io/gorm"
)

type Socmed struct {
	ID        string         `gorm:"primary_key" json:"id"`
	Title     string         `gorm:"type:varchar(255)" json:"title"`
	Content   string         `gorm:"type:varchar(255)" json:"content"`
	Tags      string         `gorm:"type:varchar(255)" json:"tags"`
	UserID    string         `gorm:"type:varchar(36);index" json:"user_id"`
	CreatedBy string         `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string         `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string         `gorm:"deleted_by,omitempty" json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
