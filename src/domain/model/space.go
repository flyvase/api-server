package model

import (
	"harvest/src/domain/value/core"
	"harvest/src/domain/value/space"
)

type Space struct {
	Id               space.Id
	Headline         string
	Access           core.Access
	NumberOfVisitors space.NumberOfVisitors
	CustomerSegment  space.CustomerSegment
	Price            space.Price
	WebsiteUrl       string
	Coordinate       core.GeoPoint
}
