package entity

import "time"

type Part struct {
	uuid         string
	name         string
	description  string
	price        float64
	category     string
	dimensions   Dimensions
	manufacturer Manufacturer
	tags         []string
	metadata     map[string]Value
	createdAt    time.Time
	updatedAt    time.Time
}

type Value struct {
	stringValue string
	int64Value  int64
	doubleValue float64
	boolValue   bool
}

type Dimensions struct {
	length float64
	width  float64
	height float64
	weight float64
}

type Manufacturer struct {
	name    string
	country string
	website string
}
