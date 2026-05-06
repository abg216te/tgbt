package repository

import "github.com/abg216te/tgbt/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	GetAll() ([]domain.User, error)
}
