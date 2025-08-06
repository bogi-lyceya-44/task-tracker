-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_tracker.priority AS ENUM (
    'unspecified',
    'low',
    'medium',
    'high',
    'critical'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE task_tracker.priority;
-- +goose StatementEnd
