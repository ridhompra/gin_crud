package models

import (
	"time"
)

type Books struct {
	Id          int64     `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Description string    `gorm:"type:text[]" json:"description"`
	CreatedAt   time.Time `json:"create_at"`
	Return      string    `gorm:"type:date;default:null" json:"return"`
}
type Users struct {
	Id       int64  `gorm:"primarykey"`
	Username string `gorm:"type:varchar(100)" json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
