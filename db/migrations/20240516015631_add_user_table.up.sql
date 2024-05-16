CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nip BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(255),
    role VARCHAR(10) NOT NULL,
    identity_card_scan_img TEXT,
    deleted_at TIMESTAMP,    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_id ON users(id);
CREATE INDEX IF NOT EXISTS idx_users_nip ON users(nip);
CREATE INDEX IF NOT EXISTS idx_users_name ON users(name);
CREATE INDEX IF NOT EXISTS idx_users_password_null ON users(password) WHERE password IS NULL;
CREATE INDEX IF NOT EXISTS idx_users_password_not_null ON users(password) WHERE password IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_role_it ON users(role) WHERE role = 'it';
CREATE INDEX IF NOT EXISTS idx_users_role_nurse ON users(role) WHERE role = 'nurse';
CREATE INDEX IF NOT EXISTS idx_users_deleted_at_null ON users(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_users_created_at_asc ON users(created_at ASC);
CREATE INDEX IF NOT EXISTS idx_users_created_at_desc ON users(created_at DESC);