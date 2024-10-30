package entity

import (
    "gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Name      string
	Description string
	Status    string
	CreatedAt string
	UpdatedAt string
}