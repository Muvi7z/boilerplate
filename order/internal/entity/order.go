package entity

type Order struct {
	orderUuid       string
	userUuid        string
	partUuids       []string
	totalPrice      float64
	transactionUuid string
	paymentMethod   string
	status          string
}
