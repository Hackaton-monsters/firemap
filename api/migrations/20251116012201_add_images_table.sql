-- +goose Up
-- +goose StatementBegin
CREATE TABLE image (
       id  INT AUTO_INCREMENT PRIMARY KEY COMMENT 'identifier',
       url VARCHAR(500) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
