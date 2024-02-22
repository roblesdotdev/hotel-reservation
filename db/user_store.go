package db

import (
	"context"
	"fmt"

	"github.com/roblesdotdev/hotel-reservation/types"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) []*types.User
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
