-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
	id serial primary key,
	usernames text[] not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd
