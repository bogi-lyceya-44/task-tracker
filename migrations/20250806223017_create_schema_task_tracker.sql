-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA task_tracker;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA task_tracker;
-- +goose StatementEnd
