CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    first_name TEXT,
    last_name  TEXT,
    email      TEXT NOT NULL UNIQUE
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);

CREATE TABLE credentials (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    email    TEXT UNIQUE,
    username TEXT NOT NULL,
    password TEXT NOT NULL,

    CONSTRAINT fk_credentials_user
        FOREIGN KEY (id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE INDEX idx_credentials_deleted_at ON credentials (deleted_at);