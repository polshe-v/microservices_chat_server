package chat

import (
	"context"
	"errors"
	"log"

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

var errQueryBuild = errors.New("failed to build query")

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
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("%v", err)
		return 0, errQueryBuild
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Printf("%v", err)
		return 0, errors.New("failed to create chat")
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("%v", err)
		return errQueryBuild
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("%v", err)
		return errors.New("failed to delete chat")
	}
	log.Printf("result: %v", res)
	return nil
}
