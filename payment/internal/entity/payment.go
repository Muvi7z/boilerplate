package entity

type Payment struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod string
}

const (
	UNKNOWN        = "UNKNOWN"
	CARD           = "CARD"
	SBP            = "SBP"
	CREDIT_CARD    = "CREDIT_CARD"
	INVESTOR_MONEY = "INVESTOR_MONEY"
)
