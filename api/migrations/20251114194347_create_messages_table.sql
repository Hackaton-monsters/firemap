-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT 'language identifier',
    chat_id    INT  NOT NULL,
    user_id    INT  NOT NULL DEFAULT 0,
    text       TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update date',
    FOREIGN KEY (chat_id) REFERENCES chats (id)
#     FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
