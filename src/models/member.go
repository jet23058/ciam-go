package models

import (
	"time"

	"github.com/google/uuid"
)

type Members struct {
	Id            uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	AccessToken   string    `gorm:"type:string;size:1500" json:"access_token"`
	IdToken       string    `gorm:"type:string;size:1500;unique" json:"id_token"`
	RefresthToken string    `gorm:"type:string;size:2000" json:"refresh_token"`
	IsUsed        bool      `gorm:"type:boolean;size:1" json:"is_used"`
	CreatedAt     time.Time `sql:"DEFAULT:'current_timestamp'" json:"created_at"`
	UpdatedAt     time.Time `sql:"DEFAULT:'current_timestamp'" json:"updated_at"`
}
