package repository

import (
	"github.com/riskykurniawan15/learn-grpc/database"
	"github.com/riskykurniawan15/learn-grpc/models"
)

// UserRepository handles database operations for users
type UserRepository struct{}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	return database.DB.Save(user).Error
}

// Delete deletes a user by ID
func (r *UserRepository) Delete(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
