-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS assets (
    id VARCHAR(26) PRIMARY KEY,
    partner_id VARCHAR(26) NOT NULL,
    name VARCHAR(100) NOT NULL,
    price BIGINT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(26) NOT NULL,
    updated_at DATETIME NULL,
    updated_by VARCHAR(26) NULL,
    CONSTRAINT uq_asset UNIQUE (partner_id, name),
    CONSTRAINT check_price CHECK (price > 0),
    CONSTRAINT fk_assets_created_by FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT fk_assets_updated_by FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS assets;
-- +goose StatementEnd