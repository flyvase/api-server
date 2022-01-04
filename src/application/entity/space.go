package entity

import (
	"api-server/src/core/constant"
	"api-server/src/domain/model"
	"api-server/src/domain/value"
	"strconv"
)

type Space struct {
	Id                  uint64
	Headline            string
	Access              string
	WeeklyVisitors      uint32
	MainCustomersSex    string
	MinMainCustomersAge uint8
	MaxMainCustomersAge uint8
	DailyPrice          uint32
	WebsiteUrl          string
	Latitude            float32
	Longitude           float32
}

func (s *Space) ToSpaceModel(imageEntities []*SpaceImage, displayEntities []*SpaceDisplay) *model.Space {
	c, err := strconv.Atoi(s.MainCustomersSex)
	if err != nil {
		panic(err)
	}
	sexCode := uint8(c)

	var imageModels []*model.SpaceImage
	for _, i := range imageEntities {
		imageModels = append(imageModels, i.ToSpaceImageModel())
	}

	var displayModels []*model.SpaceDisplay
	for _, d := range displayEntities {
		displayModels = append(displayModels, d.toSpaceDisplayModel())
	}

	return &model.Space{
		Id: value.SpaceId{
			Value: s.Id,
		},
		Headline: s.Headline,
		Access:   s.Access,
		NumberOfVisitors: value.NumberOfVisitors{
			Visitors: uint(s.WeeklyVisitors),
			Duration: constant.WeekDuration(),
		},
		CustomerSegment: value.CustomerSegment{
			Sex: value.Sex{
				Code: sexCode,
			},
			MinAge: s.MinMainCustomersAge,
			MaxAge: s.MaxMainCustomersAge,
		},
		Price: value.Price{
			Price:    uint(s.DailyPrice),
			Duration: constant.DayDuration(),
		},
		WebsiteUrl: s.WebsiteUrl,
		Coordinate: value.GeoPoint{
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		},
		Images:   imageModels,
		Displays: displayModels,
	}
}
