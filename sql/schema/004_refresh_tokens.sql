-- +goose Up
CREATE TABLE refresh_tokens (
    token TEXT PRIMARY KEY,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at timestamp with time zone NOT NULL,
    revoked_at timestamp with time zone
);

-- +goose Down
DROP TABLE refresh_tokens;