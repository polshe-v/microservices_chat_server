-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
	id serial primary key,
	usernames text[] not null,
	from_ text not null,
	text_ text not null,
    timestamp_ timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd
