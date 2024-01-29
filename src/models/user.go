package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStauts string

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string         `gorm:"type:string;size:100" json:"name"`
	Account   string         `gorm:"type:string;size:100;unique" json:"account"`
	Password  string         `gorm:"type:string;size:100" json:"-"`
	Email     string         `gorm:"type:string;size:255" json:"email"`
	Status    UserStauts     `gorm:"type:string;size:32;default:'active';check:status IN ('active', 'inactive', 'forbidden')" json:"status"`
	CreatedAt time.Time      `sql:"DEFAULT:'current_timestamp'" json:"createdAt"`
	UpdatedAt time.Time      `sql:"DEFAULT:'current_timestamp'" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
