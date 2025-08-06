-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY task_name_idx ON task_tracker.tasks USING gist(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX task_tracker.task_name_idx;
-- +goose StatementEnd
