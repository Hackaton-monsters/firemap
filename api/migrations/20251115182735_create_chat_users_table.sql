-- +goose Up
-- +goose StatementBegin
CREATE TABLE chat_users
(
    chat_id    INT NOT NULL COMMENT 'chat id',
    user_id    INT NOT NULL COMMENT 'user id',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'when user joined chat',

    PRIMARY KEY (chat_id, user_id),

    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
