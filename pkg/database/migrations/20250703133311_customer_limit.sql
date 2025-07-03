-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer_limits (
    id CHAR(26) PRIMARY KEY,
    user_id CHAR(26) NOT NULL,
    tenor SMALLINT NOT NULL,
    limit_amount BIGINT NOT NULL DEFAULT 0,
    used_amount BIGINT NOT NULL DEFAULT 0,
    available_amount BIGINT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_user_id_tenor UNIQUE (user_id, tenor),
    CONSTRAINT check_tenor CHECK (tenor IN (1, 2, 3, 4))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer_limits;
-- +goose StatementEnd