-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_tracker.topics (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    task_ids BIGINT[],
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_tracker.topics;
-- +goose StatementEnd
