ALTER TABLE users
    DROP CONSTRAINT IF EXISTS users_email_key;

ALTER TABLE credentials
    DROP CONSTRAINT IF EXISTS credentials_email_key;

ALTER TABLE credentials
    ADD CONSTRAINT credentials_username_key UNIQUE (username);

CREATE TABLE sessions (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    user_id BIGSERIAL,
    token TEXT UNIQUE NOT NULL,
    prev_token TEXT,
    expires_at TIMESTAMPTZ,
    user_agent TEXT NOT NULL,
    ip_address INET NOT NULL,

    CONSTRAINT fk_sessions_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE INDEX sessions_deleted_at ON sessions (deleted_at);