package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// estructura de usuario en la base de datos
type User struct {
	ID        string         `json:"id" gorm:"type:char(36);primary_key;unique;not null"`
	FirstName string         `json:"first_name" gorm:"not null type:char(50)"`
	LastName  string         `json:"last_name"  gorm:"not null type:char(50)"`
	Email     string         `json:"email"  gorm:"not null type:char(50)"`
	Phone     string         `json:"phone"  gorm:"not null type:char(30)"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"` //esto lo que permite es que se pueda hacer un borrado logico
}

// BeforeCreate se ejecuta antes de crear un registro de usuario
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
