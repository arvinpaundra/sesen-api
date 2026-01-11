BEGIN;

CREATE TYPE donation_status AS ENUM ('pending', 'completed', 'failed', 'cancelled', 'expired');

CREATE TYPE payment_method AS ENUM ('gopay', 'shopeepay', 'dana', 'qris', 'link_aja', 'other');

CREATE TABLE IF NOT EXISTS donations (
    id UUID PRIMARY KEY,
    to_user_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    status donation_status NOT NULL DEFAULT 'pending',
    payment_method payment_method NOT NULL DEFAULT 'other',
    payment_gateway_ref VARCHAR(255) UNIQUE,
    donor_name VARCHAR(100),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (to_user_id) REFERENCES users(id)
);

CREATE TYPE transaction_type AS ENUM ('credit', 'debit');

CREATE TYPE transaction_category AS ENUM ('donation', 'payout', 'other');

CREATE TABLE IF NOT EXISTS transaction_histories (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    type transaction_type NOT NULL,
    category transaction_category NOT NULL,
    amount BIGINT NOT NULL,
    balance_before BIGINT NOT NULL,
    balance_after BIGINT NOT NULL,
    reference_id VARCHAR(255) UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS webhook_events (
    id UUID PRIMARY KEY,
    provider VARCHAR(50),
    event_id VARCHAR(255) UNIQUE,
    payload JSONB NOT NULL,
    received_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;