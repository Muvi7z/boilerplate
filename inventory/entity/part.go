package entity

import "time"

type Part struct {
	Uuid         string
	Name         string
	Description  string
	Price        float64
	Category     string
	Dimensions   Dimensions
	Manufacturer Manufacturer
	Tags         []string
	Metadata     map[string]Value
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Value struct {
	stringValue string
	int64Value  int64
	doubleValue float64
	boolValue   bool
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}
