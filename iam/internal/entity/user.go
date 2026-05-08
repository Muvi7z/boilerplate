package entity

type User struct {
	uuid               string
	email              string
	login              string
	password           string
	notificationMethod []NotificationMethod
}

type NotificationMethod struct {
	providerName string
	target       string
}
