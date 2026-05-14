package user

import (
	"context"
	"fmt"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/iam/internal/repository/converter"
	entity2 "github.com/Muvi7z/boilerplate/iam/internal/repository/entity"
	"github.com/Muvi7z/boilerplate/platform/sql/transaction"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	var err, txErr error
	var result entity.User

	txErr = transaction.SqlxTransaction(ctx, r.db, func(tx *sqlx.Tx) error {
		reqUser := entity.User{
			Login: login,
		}
		result, err = r.getUserTx(ctx, reqUser, tx)

		notificationMethods, errGetNotification := r.getNotificationMethodsByUserUUIDTx(ctx, result.Uuid, tx)
		if errGetNotification != nil {
			return errGetNotification
		}

		result.NotificationMethods = notificationMethods

		return err
	})

	if txErr != nil {
		return result, txErr
	}

	return result, nil
}

func (r *Repository) getUserTx(ctx context.Context, user entity.User, tx *sqlx.Tx) (entity.User, error) {
	whereMap := map[string]any{}

	if user.Uuid != "" {
		whereMap["uuid"] = user.Uuid
	}

	if user.Login != "" {
		whereMap["login"] = user.Login
	}

	if user.Email != "" {
		whereMap["email"] = user.Email
	}

	sql, args, err := r.qb.Select("login").
		Columns("uuid", "email", "password", "login").
		From(usersTableName).
		Where(whereMap).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("error building request: %w", err)
	}

	var row entity2.User
	err = tx.GetContext(ctx, &row, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("error executing query: %w", err)
	}

	newUser := converter.UserFromRepository(row)

	return newUser, nil
}
