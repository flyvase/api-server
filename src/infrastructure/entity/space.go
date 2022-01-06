package entity

import (
	"api-server/src/core/constant"
	"api-server/src/domain/model"
	"api-server/src/domain/value"
	"database/sql"
	"strconv"
)

type Space struct {
	Id                  uint64
	Headline            string
	Access              sql.NullString
	WeeklyVisitors      sql.NullInt32
	MainCustomersSex    string
	MinMainCustomersAge sql.NullInt32
	MaxMainCustomersAge sql.NullInt32
	DailyPrice          uint32
	WebsiteUrl          sql.NullString
	Latitude            float32
	Longitude           float32
}

func (s *Space) ToSpaceModel(imageEntities []*SpaceImage, displayEntities []*SpaceDisplay) *model.Space {
	var imageModels []*model.SpaceImage
	for _, i := range imageEntities {
		imageModels = append(imageModels, i.ToSpaceImageModel())
	}

	var displayModels []*model.SpaceDisplay
	for _, d := range displayEntities {
		displayModels = append(displayModels, d.toSpaceDisplayModel())
	}

	var access string
	if s.Access.Valid {
		access = s.Access.String
	}

	var numberOfVisitors value.NumberOfVisitors
	if s.WeeklyVisitors.Valid {
		numberOfVisitors = value.NumberOfVisitors{
			Visitors: uint32(s.WeeklyVisitors.Int32),
			Duration: constant.WeekDuration(),
		}
	}

	c, err := strconv.Atoi(s.MainCustomersSex)
	if err != nil {
		panic(err)
	}
	sexCode := uint8(c)

	var customerSegment value.CustomerSegment
	if s.MinMainCustomersAge.Valid && s.MaxMainCustomersAge.Valid {
		customerSegment = value.CustomerSegment{
			Sex: value.Sex{
				Code: sexCode,
			},
			MinAge: uint8(s.MinMainCustomersAge.Int32),
			MaxAge: uint8(s.MaxMainCustomersAge.Int32),
		}
	}

	var websiteUrl string
	if s.WebsiteUrl.Valid {
		websiteUrl = s.WebsiteUrl.String
	}

	return &model.Space{
		Id: value.SpaceId{
			Value: s.Id,
		},
		Headline:         s.Headline,
		Access:           access,
		NumberOfVisitors: numberOfVisitors,
		CustomerSegment:  customerSegment,
		Price: value.Price{
			Price:    s.DailyPrice,
			Duration: constant.DayDuration(),
		},
		WebsiteUrl: websiteUrl,
		Coordinate: value.GeoPoint{
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		},
		Images:   imageModels,
		Displays: displayModels,
	}
}
