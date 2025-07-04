-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(26) PRIMARY KEY,
    user_id VARCHAR(26) NOT NULL,
    asset_id VARCHAR(26) NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    otr_price BIGINT NOT NULL,
    admin_fee BIGINT NOT NULL DEFAULT 0,
    installment_amount BIGINT NOT NULL,
    interest_amount BIGINT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(26) NOT NULL,
    updated_at DATETIME NULL,
    updated_by VARCHAR(26) NULL,
    CONSTRAINT fk_transactions_created_by FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT fk_transactions_updated_by FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd