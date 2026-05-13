package user

import (
	"context"
	"fmt"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/platform/sql/transaction"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	usersTableName               = "users"
	notificationMethodsTableName = "notification_methods"
)

func (r *Repository) Create(ctx context.Context, user *entity.User) (string, error) {
	var res string
	var err, txErr error

	txErr = transaction.SqlxTransaction(ctx, r.db, func(tx *sqlx.Tx) error {
		userUuid := uuid.New().String()

		user.Uuid = userUuid

		res, err = r.createUserTx(ctx, user, tx)
		if err != nil {
			return err
		}

		for _, nm := range user.NotificationMethods {
			err = r.createUserNotificationMethodsTx(ctx, nm, user.Uuid, tx)
			if err != nil {
				return err
			}

		}

		return err
	})

	if txErr != nil {
		return "", txErr
	}

	return res, nil
}

func (r *Repository) createUserTx(ctx context.Context, user *entity.User, tx *sqlx.Tx) (string, error) {
	insertMap := map[string]interface{}{
		"uuid":     user.Uuid,
		"login":    user.Login,
		"email":    user.Email,
		"password": user.Password,
	}

	sql, args, err := r.qb.Insert(usersTableName).
		SetMap(insertMap).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("error building insert query: %w", err)
	}

	err = tx.GetContext(ctx, nil, sql, args...)
	if err != nil {
		return "", fmt.Errorf("error executing insert query: %w", err)
	}

	return user.Uuid, nil
}

func (r *Repository) createUserNotificationMethodsTx(ctx context.Context, notification entity.NotificationMethod, userUUID string, tx *sqlx.Tx) error {
	insertMap := map[string]interface{}{
		"user_uuid":     userUUID,
		"provider_name": notification.ProviderName,
		"target":        notification.Target,
	}

	sql, args, err := r.qb.Insert(notificationMethodsTableName).
		SetMap(insertMap).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building insert query: %w", err)
	}

	return tx.GetContext(ctx, nil, sql, args...)
}
