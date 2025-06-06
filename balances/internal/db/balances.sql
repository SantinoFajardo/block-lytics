CREATE TABLE IF NOT EXISTS balances (
    id SERIAL PRIMARY KEY,
    account_address VARCHAR(255) NOT NULL,
    blockchain VARCHAR(255) NOT NULL,
    token_address VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);