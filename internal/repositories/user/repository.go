package user

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	DefaultRepository struct {
		sqlx *sqlx.DB
		tx   *trmsqlx.CtxGetter
	}

	sqlxUserSource struct {
		UserIDint uuid.UUID `db:"user_id_int"`
		UserIDext string    `db:"user_id_ext"`
		Source    string    `db:"source"`
	}
)

func NewRepository(sqlx *sqlx.DB, tx *trmsqlx.CtxGetter) *DefaultRepository {
	return &DefaultRepository{sqlx: sqlx, tx: tx}
}

func (r *DefaultRepository) db(ctx context.Context) trmsqlx.Tr {
	return r.tx.DefaultTrOrDB(ctx, r.sqlx)
}

func (r *DefaultRepository) InsertSource(ctx context.Context, user model.UserSource) error {
	const query = `
	   insert into user_source
	   (user_id_int, user_id_ext, "source")
	   values ($1, $2, $3)
	   on conflict (user_id_ext, "source") do nothing
	`

	_, err := r.db(ctx).ExecContext(
		ctx,
		query,
		user.ID,
		user.IDext,
		user.Source,
	)

	return err
}

func (r *DefaultRepository) GetUserBySource(ctx context.Context, userIDext string, source string) (model.User, error) {
	const query = `
  	   select user_id_int, user_id_ext, "source"
  	   from user_source
  	   where user_id_ext= $1 and source = $2
	`

	var result []sqlxUserSource
	err := r.db(ctx).SelectContext(ctx, &result, query, userIDext, source)
	if err != nil {
		return model.User{}, err
	}
	if len(result) == 0 {
		return model.User{}, contracts.ErrUserNotFound
	}

	return model.User{
		ID: result[0].UserIDint,
	}, nil
}
