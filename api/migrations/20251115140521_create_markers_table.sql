-- +goose Up
-- +goose StatementBegin
CREATE TABLE markers
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT 'identifier',
    chat_id    INT           NULL COMMENT 'related chat id',
    lat        DECIMAL(9, 6) NOT NULL COMMENT 'latitude',
    lon        DECIMAL(9, 6) NOT NULL COMMENT 'longitude',
    type       VARCHAR(100)  NOT NULL COMMENT 'marker type',
    title      VARCHAR(255)  NOT NULL COMMENT 'title',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update date',
    FOREIGN KEY (chat_id) REFERENCES chats (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
