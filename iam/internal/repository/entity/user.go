package entity

type User struct {
	Uuid     string `db:"uuid"`
	Login    string `db:"login"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type NotificationMethod struct {
	Uuid         string `db:"uuid"`
	UserUUID     string `db:"user_uuid"`
	ProviderName string `db:"provider_name"`
	Target       string `db:"target"`
}
