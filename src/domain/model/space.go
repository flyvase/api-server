package model

import (
	"harvest/src/domain/value/core"
	"harvest/src/domain/value/space"
)

type Space struct {
	Id               space.Id
	Headline         string
	Access           string
	NumberOfVisitors space.NumberOfVisitors
	CustomerSegment  space.CustomerSegment
	Price            space.Price
	WebsiteUrl       string
	Coordinate       core.GeoPoint
	Images           []*SpaceImage
	Displayers       []*SpaceDisplayer
}
