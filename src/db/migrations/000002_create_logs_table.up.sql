CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    success BOOLEAN NOT NULL,
    operation_code VARCHAR(255) NOT NULL,
    ip VARCHAR(30) NOT NULL,
    timestamp TIMESTAMP NOT NULL
);