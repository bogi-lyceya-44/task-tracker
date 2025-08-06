-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY board_name_idx ON task_tracker.boards USING gist(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX task_tracker.board_name_idx;
-- +goose StatementEnd
