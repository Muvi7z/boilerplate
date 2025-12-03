package entity

import (
	"errors"
	"time"
)

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderIsPaid = errors.New("order is paid")
var ErrCreateOrder = errors.New("error creating order")
var ErrGetOrder = errors.New("error getting order")
var ErrUpdateOrder = errors.New("error updating order")

type Order struct {
	OrderUuid       string
	UserUuid        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUuid string
	PaymentMethod   string
	Status          string
}

type PayOrder struct {
	OrderUuid     string
	UserUuid      string
	PaymentMethod string
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []string
	ManufacturerCountries []string
	Tags                  []string
}

type CreateOrder struct {
	UserUuid  string
	PartUuids []string
}

const (
	UNKNOWN       = "UNKNOWN"
	CARD          = "CARD"
	SBP           = "SBP"
	CreditCard    = "CREDIT_CARD"
	InvestorMoney = "INVESTOR_MONEY"
)

const (
	CATEGORY_UNKNOWN = "UNKNOWN"
	CategoryEngine   = "ENGINE"
	CategoryFuel     = "FUEL"
	CategoryPorthole = "PORTHOLE"
	CategoryWing     = "WING"
)

type DimensionsInfo struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type ManufacturerInfo struct {
	Name    string
	Country string
	Website string
}

type Part struct {
	Uuid         string
	Name         string
	Description  string
	Price        float64
	Category     string
	Dimensions   DimensionsInfo
	Manufacturer ManufacturerInfo
	Tags         []string
	Metadata     map[string]interface{}
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
