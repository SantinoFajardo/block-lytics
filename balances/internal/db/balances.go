package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type BalancesSchema struct {
	Id        int64     `json:"id" db:"id"`
	AccountAddress string    `json:"account_address" db:"account_address"`
	Blockchain     string    `json:"blockchain" db:"blockchain"`
	TokenAddress   string    `json:"token_address" db:"token_address"`
	Amount         int64     `json:"amount" db:"amount"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type BalancesRepository interface {
	UpdateUserBalances(accountAddress string, blockchain string, tokenAddress string, amount int64) (bool, error)
}

type repository struct {
	db *sqlx.DB
}

func NewBalancesRepository(db *sqlx.DB) BalancesRepository {
	return &repository{db: db}
}

func (r *repository) UpdateUserBalances(accountAddress string, blockchain string, tokenAddress string, amount int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO balances (account_address, blockchain, token_address, amount)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (account_address, blockchain, token_address)
	DO UPDATE SET amount = balances.amount + $4
	RETURNING id
	`

	var id int64
	err := r.db.QueryRowxContext(ctx, query, accountAddress, blockchain, tokenAddress, amount).Scan(&id)
	if err != nil {
		return false, err
	}

	return true, nil
}

