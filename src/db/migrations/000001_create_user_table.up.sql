CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN,
    inactive BOOLEAN,
    verified_email BOOLEAN,

    code_verify_email VARCHAR(12),
    code_verify_email_expiry TIMESTAMP WITH TIME ZONE,

    code_change_email VARCHAR(12),
    code_change_email_expiry TIMESTAMP WITH TIME ZONE,

    password_recovery_code VARCHAR(12),
    password_recovery_code_expiry TIMESTAMP WITH TIME ZONE,

    twofa BOOLEAN,
    twofa_secret VARCHAR(255)
);