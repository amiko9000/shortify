package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        	uuid.UUID		`gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt 	time.Time		`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time		`gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}

func (m *Model) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}