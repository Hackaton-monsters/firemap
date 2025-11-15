-- +goose Up
-- +goose StatementBegin
ALTER TABLE reports
    MODIFY COLUMN comment VARCHAR(300)
        CHARACTER SET utf8mb4
        COLLATE utf8mb4_unicode_ci
        NULL COMMENT 'comment';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
