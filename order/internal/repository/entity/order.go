package entity

import "github.com/lib/pq"

type Order struct {
	OrderUuid       string         `db:"order_uuid"`
	UserUuid        string         `db:"user_uuid"`
	PartUuids       pq.StringArray `db:"part_uuids"`
	TotalPrice      float64        `db:"total_price"`
	TransactionUuid *string        `db:"transaction_uuid"`
	PaymentMethod   *string        `db:"payment_method"`
	Status          string         `db:"status"`
}
