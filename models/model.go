package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model        // Ini akan menambahkan kolom seperti ID, CreatedAt, dan UpdatedAt secara otomatis.
	Username   string `gorm:"uniqueIndex"`
	Password   string
	Email      string `gorm:"uniqueIndex"`
	FullName   string
	Role       string
}
