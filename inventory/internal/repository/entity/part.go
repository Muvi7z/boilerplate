package entity

import (
	"time"
)

type DimensionsInfo struct {
	Length float64 `bson:"length"`
	Width  float64 `bson:"width"`
	Height float64 `bson:"height"`
	Weight float64 `bson:"weight"`
}

type ManufacturerInfo struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website"`
}

type Part struct {
	Uuid         string                 `bson:"_id"`
	Name         string                 `bson:"name"`
	Description  string                 `bson:"description"`
	Price        float64                `bson:"price"`
	Category     string                 `bson:"category"`
	Dimensions   DimensionsInfo         `bson:"dimensions"`
	Manufacturer ManufacturerInfo       `bson:"manufacturer"`
	Tags         []string               `bson:"tags"`
	Metadata     map[string]interface{} `bson:"metadata"`
	CreatedAt    *time.Time             `bson:"created_at"`
	UpdatedAt    *time.Time             `bson:"updated_at"`
}
