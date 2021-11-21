package entity

import (
	"harvest/src/core/constants"
	"harvest/src/domain/model"
	"harvest/src/domain/value"
	"strconv"
)

type Space struct {
	Id                  uint32
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

func (s *Space) ToSpaceModel(imageEntities []*SpaceImage, displayerEntities []*SpaceDisplayer) *model.Space {
	c, err := strconv.Atoi(s.MainCustomersSex)
	if err != nil {
		panic(err)
	}
	sexCode := uint8(c)

	var imageModels []*model.SpaceImage
	for _, i := range imageEntities {
		imageModels = append(imageModels, i.toSpaceImageModel())
	}

	var displayerModels []*model.SpaceDisplayer
	for _, d := range displayerEntities {
		displayerModels = append(displayerModels, d.toSpaceDisplayerModel())
	}

	return &model.Space{
		Id: value.SpaceId{
			Value: uint(s.Id),
		},
		Headline: s.Headline,
		Access:   s.Access,
		NumberOfVisitors: value.NumberOfVisitors{
			Visitors: uint(s.WeeklyVisitors),
			Duration: constants.WeekDuration(),
		},
		CustomerSegment: value.CustomerSegment{
			Sex: value.Sex{
				Code: sexCode,
			},
			MinAge: s.MinMainCustomersAge,
			MaxAge: s.MaxMainCustomersAge,
		},
		Price: value.Price{
			Value:    uint(s.DailyPrice),
			Duration: constants.DayDuration(),
		},
		WebsiteUrl: s.WebsiteUrl,
		Coordinate: value.GeoPoint{
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		},
		Images:     imageModels,
		Displayers: displayerModels,
	}
}
