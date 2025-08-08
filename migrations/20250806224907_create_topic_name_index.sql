-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY topic_name_idx ON task_tracker.topics USING gist(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX task_tracker.topic_name_idx;
-- +goose StatementEnd
