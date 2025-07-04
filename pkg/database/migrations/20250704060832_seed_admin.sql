-- +goose Up
-- +goose StatementBegin
INSERT INTO "users"
(id, nik, full_name, legal_name, place_of_birth, date_of_birth, salary, id_card_photo_url, selfie_photo_url, password, is_admin, created_at, updated_at)
VALUES('01JZ86DBY6NXRZQ3XHE4XBN5CZ', '1111111111111111', 'John Doe', 'Johnathan Doe', 'Cityville', '1990-01-01', 5000000, 'http://example.com/id_card.jpg', 'http://example.com/selfie.jpg', '$2a$10$ogyNLaAPP5.Md119CVm8ru/pxT63CdCDvo8gM/HYZa0dMy94z6YYq', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); -- password: Password123!@#
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "users" WHERE id = '01JZ86DBY6NXRZQ3XHE4XBN5CZ';
-- +goose StatementEnd
