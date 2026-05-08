package entity

type SessionRedisView struct {
	Uuid   string `redis:"uuid"`
	UserId string `redis:"user_id"`
}
