CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    name VARCHAR(50) NOT NULL,
    description TEXT,
    status VARCHAR(20),
    board_id BIGINT NOT NULL,

    CONSTRAINT fk_tasks_board
        FOREIGN KEY (board_id)
        REFERENCES boards (id)
        ON DELETE CASCADE
);

CREATE INDEX idx_tasks_board_id ON tasks (board_id);
CREATE INDEX idx_tasks_deleted_at ON tasks (deleted_at);
