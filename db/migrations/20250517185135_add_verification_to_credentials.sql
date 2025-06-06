-- +goose Up
-- +goose StatementBegin
ALTER TABLE credentials ADD COLUMN is_verification BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE credentials DROP COLUMN is_verification;
-- +goose StatementEnd

