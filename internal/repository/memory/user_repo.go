package memory

import (
	"sync"

	"github.com/abg216te/tgbt/internal/domain"
)

type UserRepo struct {
	mu    sync.Mutex
	users []domain.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users = append(r.users, *user)
	return nil
}

func (r *UserRepo) GetAll() ([]domain.User, error) {
	return r.users, nil
}
