package service

import (
	"github.com/abg216te/tgbt/internal/domain"
	"github.com/abg216te/tgbt/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(tgID int64, name string) error {
	user := &domain.User{
		TelegramID: tgID,
		Name:       name,
	}
	return s.repo.Save(user)
}

func (s *UserService) List() ([]*domain.User, error) {
	return s.repo.GetAll()
}
