package db

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountSchema struct {
	Id        int64     `json:"id" db:"id"`
	Address   string    `json:"address" db:"address"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Blockchain string    `json:"blockchain" db:"blockchain"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AccountRepository interface {
	GetByAddress(address string) (*AccountSchema, error)
	Create(address string, user_id int64, blockchain string) (int64, error)
}

type repository struct {
	db *sqlx.DB
}	

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &repository{db: db}
}

func (r *repository) GetByAddress(address string) (*AccountSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, address, user_id, blockchain, created_at, updated_at
		FROM accounts
		WHERE address = $1
		`

	var account AccountSchema
	row := r.db.QueryRowContext(ctx, query, address)
	err := row.Scan(&account.Id, &account.Address, &account.UserId,&account.Blockchain, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return &account, nil
}

func (r *repository) Create(address string, user_id int64, blockchain string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Creating account with address:", address, "and user_id:", user_id, "and blockchain:", blockchain)

	query := `
		INSERT INTO accounts (address, user_id, blockchain) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	err := r.db.GetContext(ctx, &id, query, address, user_id, blockchain)
	if err != nil {
		return 0, err
	}
	return id, nil
}