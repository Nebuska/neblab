CREATE TABLE board_users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    user_id BIGINT NOT NULL,
    board_id BIGINT NOT NULL,
    role VARCHAR(255),

    CONSTRAINT fk_board_users_board
        FOREIGN KEY (board_id)
        REFERENCES boards (id)
        ON DELETE CASCADE
);

CREATE INDEX idx_board_users_user_id ON board_users (user_id);
CREATE INDEX idx_board_users_board_id ON board_users (board_id);
CREATE INDEX idx_board_users_deleted_at ON board_users (deleted_at);

CREATE UNIQUE INDEX uniq_board_user
    ON board_users (user_id, board_id)
    WHERE deleted_at IS NULL;