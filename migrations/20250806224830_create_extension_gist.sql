-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION btree_gist;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION btree_gist;
-- +goose StatementEnd
