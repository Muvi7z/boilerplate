package entity

type User struct {
	uuid     string `db:"uuid"`
	login    string `db:"login"`
	email    string `db:"email"`
	password string `db:"password"`
}

type NotificationMethod struct {
	uuid         string `db:"uuid"`
	userUUID     string `db:"user_uuid"`
	providerName string `db:"provider_name"`
	target       string `db:"target"`
}
