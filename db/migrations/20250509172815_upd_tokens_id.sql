-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE issued_jwt_token
    DROP CONSTRAINT issued_jwt_token_pkey;

ALTER TABLE issued_jwt_token
    ADD COLUMN id SERIAL;

ALTER TABLE issued_jwt_token
    ADD CONSTRAINT issued_jwt_token_pkey PRIMARY KEY (id);

ALTER TABLE issued_jwt_token
    ADD CONSTRAINT issued_jwt_token_jti_key UNIQUE (jti);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE issued_jwt_token
    DROP CONSTRAINT issued_jwt_token_jti_key;

ALTER TABLE issued_jwt_token
    DROP CONSTRAINT issued_jwt_token_pkey;

ALTER TABLE issued_jwt_token
    DROP COLUMN id;

ALTER TABLE issued_jwt_token
    ADD CONSTRAINT issued_jwt_token_pkey PRIMARY KEY (jti);
-- +goose StatementEnd
