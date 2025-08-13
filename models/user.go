package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user table in database
type User struct {
	ID        uint           `gorm:"primarykey" json:"id" validate:"-"`
	Name      string         `gorm:"size:100;not null" json:"name" validate:"required,min=2,max=100,alpha_space"`
	Email     string         `gorm:"size:100;unique;not null" json:"email" validate:"required,email,max=100"`
	Password  string         `gorm:"size:255;not null" json:"-" validate:"required,min=8,max=255,password_strength"`
	Age       int            `gorm:"not null" json:"age" validate:"required,min=13,max=120"`
	CreatedAt time.Time      `json:"created_at" validate:"-"`
	UpdatedAt time.Time      `json:"updated_at" validate:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" validate:"-"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

// Validation request structs
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Age      int    `json:"age" validate:"required,min=13,max=120"`
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=2,max=100,alpha_space"`
	Email    string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	Password string `json:"password,omitempty" validate:"omitempty,min=8,max=255,password_strength"`
	Age      int    `json:"age,omitempty" validate:"omitempty,min=13,max=120"`
}
