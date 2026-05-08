package user

import (
	"context"
	"os/user"
)

func (r *repository) getUserByID(ctx context.Context, uuid string) (user.User, error) {
	whereMap := map[string]any{
		"uuid": uuid,
	}

	r.qb.Select("uuid").
		Columns("uuid", "email", "password", "login")

	return user.User{}
}
