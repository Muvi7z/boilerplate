package entity

type User struct {
	Uuid                string
	Email               string
	Login               string
	Password            string
	NotificationMethods []NotificationMethod
}

type NotificationMethod struct {
	ProviderName string
	Target       string
}
