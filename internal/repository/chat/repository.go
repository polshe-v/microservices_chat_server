package chat

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/polshe-v/microservices_chat_server/internal/repository"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

const (
	tableName = "chats"

	idColumn        = "id"
	usernamesColumn = "usernames"
)

var errQueryBuild = errors.New("failed to build query")

type repo struct {
	db *pgxpool.Pool
}

// NewRepository creates new object of repository layer.
func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *desc.Chat) (int64, error) {
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

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
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

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("%v", err)
		return errors.New("failed to delete chat")
	}
	log.Printf("result: %v", res)
	return nil
}
