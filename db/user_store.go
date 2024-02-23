package db

import (
	"context"
	"fmt"

	"github.com/roblesdotdev/hotel-reservation/types"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) []*types.User
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
}

type MemoryUserStore struct {
	store []*types.User
}

func NewMemoryUserStore(store []*types.User) *MemoryUserStore {
	return &MemoryUserStore{
		store: store,
	}
}

func (s *MemoryUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var u *types.User
	for _, user := range s.store {
		if user.ID == id {
			u = user
		}
	}
	if u == nil {
		return nil, fmt.Errorf("could not found user with id %s", id)
	}
	return u, nil
}

func (s *MemoryUserStore) GetAllUsers(ctx context.Context) []*types.User {
	return s.store
}

func (s *MemoryUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	s.store = append(s.store, user)
	return user, nil
}

func (s *MemoryUserStore) DeleteUser(ctx context.Context, id string) error {
	for i, user := range s.store {
		if user.ID == id {
			s.store = append(s.store[:i], s.store[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not found user with id %s", id)
}

func (s *MemoryUserStore) UpdateUser(ctx context.Context, id string, updates *types.User) error {
	return nil
}
