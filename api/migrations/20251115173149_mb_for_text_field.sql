-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages
    MODIFY COLUMN text TEXT
        CHARACTER SET utf8mb4
        COLLATE utf8mb4_unicode_ci
        NOT NULL COMMENT 'message text';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
