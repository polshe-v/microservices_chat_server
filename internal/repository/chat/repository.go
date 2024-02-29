package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/polshe-v/microservices_chat_server/internal/client/db"
	"github.com/polshe-v/microservices_chat_server/internal/model"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
)

const (
	tableName = "chats"

	idColumn        = "id"
	usernamesColumn = "usernames"
)

type repo struct {
	db db.Client
}

// NewRepository creates new object of repository layer.
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesColumn).
		Values(chat.Usernames).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
