package store

import "github.com/RomanKazakevich/group-materials/internal/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(user *model.User) (*model.User, error) {

}
