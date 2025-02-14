# Chirpy

Chirpy is a simple social media application where users can post short messages called chirps. This project is built using Go and PostgreSQL.

## Features

- User authentication (login, refresh, revoke)
- Create, read, and delete chirps
- Sort chirps by creation date
- Admin metrics and reset functionality
- Webhook handling for user upgrades

## Endpoints

### Authentication

- `POST /api/login`: User login
- `POST /api/refresh`: Refresh JWT token
- `POST /api/revoke`: Revoke JWT token

### Chirps

- `POST /api/chirps`: Create a new chirp
- `GET /api/chirps`: Get all chirps (supports sorting by `created_at` with `sort` query parameter)
- `GET /api/chirps/{chirpID}`: Get a specific chirp by ID
- `DELETE /api/chirps/{chirpID}`: Delete a specific chirp by ID

### Users

- `POST /api/users`: Create a new user
- `PUT /api/users`: Update user information

### Admin

- `GET /admin/metrics`: Get admin metrics
- `POST /admin/reset`: Reset admin metrics

### Health Check

- `GET /api/healthz`: Health check endpoint

### Webhooks

- `POST /api/polka/webhooks`: Handle Polka webhooks for user upgrades

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/eefret/chirpy.git
    cd chirpy
    ```

2. Create a [.env](http://_vscodecontentref_/1) file with the following environment variables:
    ```env
    DB_URL=your_database_url
    AUTH_SECRET=your_auth_secret
    POLKA_KEY=your_polka_key
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

4. The server will start on `http://localhost:8080`.

## Running Tests

To run the tests, use the following command:
```sh
go test ./...