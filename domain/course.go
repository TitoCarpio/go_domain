package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string         `json:"id"  gorm:"type:char(36);primary_key;unique;not null"`
	Name      string         `json:"name" gorm:"not null type:char(50)"`
	StartDate time.Time      `json:"start_date" `
	EndDate   time.Time      `json:"ends_date" `
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

// BeforeCreate se ejecuta antes de crear un registro de curso
func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
