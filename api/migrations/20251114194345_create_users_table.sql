-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT 'identifier',
    email      VARCHAR(1000) NOT NULL COMMENT 'email',
    password   VARCHAR(1000) NOT NULL COMMENT 'password',
    nickname   VARCHAR(1000) NOT NULL COMMENT 'nickname',
    role       VARCHAR(100)  NOT NULL COMMENT 'user role',
    token      VARCHAR(1000) COMMENT 'token',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update date',
    UNIQUE KEY uq_users_email (email),
    UNIQUE KEY uq_users_nickname (nickname)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
