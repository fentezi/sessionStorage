package service

import (
	"errors"

	"github.com/fentezi/session-auth/internal/models"
	repostirories "github.com/fentezi/session-auth/internal/repositories"
	"github.com/fentezi/session-auth/pkg"
)

type Service struct {
	repo *repostirories.Repositories
}

func NewService(repo *repostirories.Repositories) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(user *models.User) (uint, error) {
	hashPass, err := pkg.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashPass
	userID, err := s.repo.PostgreSQL.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *Service) SignIn(user *models.User) (string, error) {
	dbUser, err := s.repo.PostgreSQL.GetUser(user.Email)
	if err != nil {
		return "", err
	}
	coincidence := pkg.CheckPasswordHash(user.Password, dbUser.Password)
	if coincidence {
		uuid := pkg.GenerateUUID()
		err := s.repo.Redis.CreateSession(uuid, dbUser.ID)
		if err != nil {
			return "", err
		}
		return uuid, nil
	}
	return "", errors.New("неверный пароль или логин")
}
