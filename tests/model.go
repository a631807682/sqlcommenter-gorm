package tests

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint `gorm:"default:18"`
	Birthday *time.Time
	Active   bool
}
