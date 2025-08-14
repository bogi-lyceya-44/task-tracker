-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_tracker.board_order (
    id BIGINT PRIMARY KEY REFERENCES task_tracker.boards(id),
    place INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_tracker.board_order;
-- +goose StatementEnd
