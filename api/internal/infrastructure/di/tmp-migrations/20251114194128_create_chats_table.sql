-- +goose Up
-- +goose StatementBegin

CREATE TABLE chats
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT 'language identifier',
    name       VARCHAR(1000) NOT NULL COMMENT 'language name',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update date'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
