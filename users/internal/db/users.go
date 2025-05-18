package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserSchema struct {
	Id        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserRepository interface {
	CreateUser(email string,password string)(int64,error)
}

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) CreateUser(email string, password string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id
	`

	var id int64
	err := r.db.QueryRowxContext(ctx, query, email, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
