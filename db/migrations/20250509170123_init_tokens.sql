-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE issued_jwt_token (
    jti VARCHAR(36) PRIMARY KEY,                    -- Уникальный идентификатор токена (JWT ID)
    subject_id INTEGER NOT NULL,                    -- Внешний ключ на пользователя
    token_type INTEGER NOT NULL,                -- Тип токена: 'access' или 'refresh'
    revoked BOOLEAN DEFAULT FALSE,                  -- Флаг, указывающий на то, был ли токен отозван
    issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Время выпуска токена
    expires_at TIMESTAMP,                           -- Время истечения токена
    CONSTRAINT fk_user FOREIGN KEY (subject_id) REFERENCES credentials (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE issued_jwt_token;
-- +goose StatementEnd
