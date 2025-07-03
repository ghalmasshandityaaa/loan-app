-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
    id CHAR(26) PRIMARY KEY,
    nik VARCHAR(16) NOT NULL UNIQUE, -- National ID number
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100) NOT NULL, -- Name as shown on the ID
    place_of_birth VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    salary BIGINT NOT NULL,
    id_card_photo_url TEXT NOT NULL,
    selfie_photo_url TEXT NOT NULL,
    password text NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
