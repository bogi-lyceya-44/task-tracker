-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY task_priority_idx ON task_tracker.tasks(priority);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX task_tracker.task_priority_idx;
-- +goose StatementEnd
