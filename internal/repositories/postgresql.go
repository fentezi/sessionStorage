package repostirories

import (
	"github.com/fentezi/session-auth/config"
	"github.com/fentezi/session-auth/internal/models"
)

type PostgreSQL struct {
}

func (p *PostgreSQL) CreateUser(user *models.User) (uint, error) {
	result := config.DB.Create(user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (p *PostgreSQL) GetUser(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repositories) DeleteUser(email string) error {
	result := config.DB.Where("email = ?", email).Delete(&models.User{})
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
