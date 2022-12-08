package store

import "github.com/beeloon/go-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	createQuery := "INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id"
	if err := r.store.db.QueryRow(createQuery, u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	findByEmailQuery := "SELECT id, email, encrypted_password FROM users WHERE email = $1"
	if err := r.store.db.QueryRow(findByEmailQuery, email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
