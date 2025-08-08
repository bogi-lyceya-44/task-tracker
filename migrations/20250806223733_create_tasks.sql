-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_tracker.tasks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    dependencies BIGINT[],
    priority INT NOT NULL,
    start_time TIMESTAMP,
    finish_time TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_tracker.tasks;
-- +goose StatementEnd
