-- +goose Up
CREATE TABLE chirps (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body TEXT NOT NULL
);

-- +goose Down
DROP TABLE chirps;