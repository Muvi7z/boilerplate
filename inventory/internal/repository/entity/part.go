package entity

import "time"

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
