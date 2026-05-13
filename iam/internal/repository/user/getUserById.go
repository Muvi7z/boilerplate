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

func (r *Repository) Get(ctx context.Context, uuid string) (entity.User, error) {
	var err, txErr error
	var result entity.User

	txErr = transaction.SqlxTransaction(ctx, r.db, func(tx *sqlx.Tx) error {
		result, err = r.getUserByIDTx(ctx, uuid, tx)

		notificationMethods, errGetNotification := r.getNotificationMethodsByUserUUIDTx(ctx, uuid, tx)
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

func (r *Repository) getUserByIDTx(ctx context.Context, uuid string, tx *sqlx.Tx) (entity.User, error) {
	whereMap := map[string]any{
		"uuid": uuid,
	}

	sql, args, err := r.qb.Select("uuid").
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

	user := converter.UserFromRepository(row)

	return user, nil
}

func (r *Repository) getNotificationMethodsByUserUUIDTx(ctx context.Context, userUuid string, tx *sqlx.Tx) ([]entity.NotificationMethod, error) {
	whereMap := map[string]any{
		"user_uuid": userUuid,
	}

	sql, args, err := r.qb.Select("user_uuid").
		Columns("uuid", "user_uuid", "provider_name", "target").
		From(notificationMethodsTableName).
		Where(whereMap).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	var rows []entity2.NotificationMethod
	err = tx.SelectContext(ctx, &rows, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	var result []entity.NotificationMethod
	for _, row := range rows {
		result = append(result, converter.NotificationMethodFromRepository(row))
	}

	return result, nil
}
