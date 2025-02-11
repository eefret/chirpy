-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    email TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;