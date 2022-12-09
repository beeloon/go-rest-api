package teststore

import (
	"github.com/beeloon/go-rest-api/internal/app/model"
	"github.com/beeloon/go-rest-api/internal/app/store"
	_ "github.com/lib/pq" // ...
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}

	return s.userRepository
}
