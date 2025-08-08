-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_tracker.boards (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    topic_ids BIGINT[],
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_tracker.boards;
-- +goose StatementEnd
