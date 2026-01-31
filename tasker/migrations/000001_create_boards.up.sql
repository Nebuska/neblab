CREATE TABLE boards (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    name VARCHAR(30) NOT NULL
);

CREATE INDEX idx_boards_deleted_at ON boards (deleted_at);