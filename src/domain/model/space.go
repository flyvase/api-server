package model

import (
	"harvest/src/domain/value"
)

type Space struct {
	Id               value.SpaceId
	Headline         string
	Access           string
	NumberOfVisitors value.NumberOfVisitors
	CustomerSegment  value.CustomerSegment
	Price            value.Price
	WebsiteUrl       string
	Coordinate       value.GeoPoint
	Images           []*SpaceImage
	Displayers       []*SpaceDisplayer
}
