-- +goose Up
-- +goose StatementBegin
CREATE TABLE reports
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT 'identifier',
    comment    VARCHAR(300),
    photos     VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'keywords for meta',
    marker_id  INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update date',
    FOREIGN KEY (marker_id) REFERENCES markers (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
