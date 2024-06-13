// Only used for testing. Will be removed later

package persistence

import (
	"errors"
	"tometower/internal/domain/user"
)

type InMemoryRepository struct {
	data map[string]user.User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[string]user.User)}
}

func (r *InMemoryRepository) Save(entity user.User) error {
	r.data[entity.ID] = entity
	return nil
}

func (r *InMemoryRepository) FindByID(id string) (user.User, error) {
	entity, exists := r.data[id]
	if !exists {
		return user.User{}, errors.New("user not found")
	}
	return entity, nil
}
