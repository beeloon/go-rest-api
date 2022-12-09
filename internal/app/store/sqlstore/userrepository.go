package sqlstore

import (
	"database/sql"

	"github.com/beeloon/go-rest-api/internal/app/model"
	"github.com/beeloon/go-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	query := "INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id"
	return r.store.db.QueryRow(query, u.Email, u.EncryptedPassword).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	query := "SELECT id, email, encrypted_password FROM users WHERE email = $1"
	if err := r.store.db.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
