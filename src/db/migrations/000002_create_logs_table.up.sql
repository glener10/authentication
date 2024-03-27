CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    find_param VARCHAR(60) NOT NULL,
    route VARCHAR(50) NOT NULL,
    method VARCHAR(15) NOT NULL,
    success BOOLEAN NOT NULL,
    operation_code VARCHAR(100) NOT NULL,
    ip VARCHAR(30) NOT NULL,
    timestamp TIMESTAMP NOT NULL
);