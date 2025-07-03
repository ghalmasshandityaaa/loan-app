-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS partners (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    partner_type ENUM('ecommerce', 'dealer') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(26) NOT NULL,
    updated_at DATETIME NULL,
    updated_by VARCHAR(26) NULL,
    CONSTRAINT uq_partner_name UNIQUE (name),
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS partners;
-- +goose StatementEnd
