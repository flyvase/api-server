package entity

import (
	"harvest/src/domain/object"
	"time"
)

type SpaceImage struct {
	Id        uint32
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type SpaceDisplayer struct {
	Id          uint32
	ImageUrl    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Space struct {
	Id                  uint32
	Headline            string
	Access              string
	NumberOfVisitors    uint32
	MainCustomersSex    string
	MinMainCustomersAge uint8
	MaxMainCustomersAge uint8
	Price               uint32
	WebsiteUrl          string
	Coordinate          object.GeoPoint
	Images              []*SpaceImage
	Displayers          []*SpaceDisplayer
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}

type SpaceOwner struct {
	Id        uint32
	Name      string
	Spaces    []*Space
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
