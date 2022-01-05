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

	var convertedAccess string
	if s.Access.Valid {
		convertedAccess = s.Access.String
	}

	var convertedWeeklyVisitors uint32
	if s.WeeklyVisitors.Valid {
		convertedWeeklyVisitors = uint32(s.WeeklyVisitors.Int32)
	}

	var convertedMinMainCustomersAge uint8
	if s.MinMainCustomersAge.Valid {
		convertedMinMainCustomersAge = uint8(s.MinMainCustomersAge.Int32)
	}

	var convertedMaxMainCustomersAge uint8
	if s.MaxMainCustomersAge.Valid {
		convertedMaxMainCustomersAge = uint8(s.MaxMainCustomersAge.Int32)
	}

	var convertedWebsiteUrl string
	if s.WebsiteUrl.Valid {
		convertedWebsiteUrl = s.WebsiteUrl.String
	}

	return &model.Space{
		Id: value.SpaceId{
			Value: s.Id,
		},
		Headline: s.Headline,
		Access:   convertedAccess,
		NumberOfVisitors: value.NumberOfVisitors{
			Visitors: uint(convertedWeeklyVisitors),
			Duration: constant.WeekDuration(),
		},
		CustomerSegment: value.CustomerSegment{
			Sex: value.Sex{
				Code: sexCode,
			},
			MinAge: convertedMinMainCustomersAge,
			MaxAge: convertedMaxMainCustomersAge,
		},
		Price: value.Price{
			Price:    uint(s.DailyPrice),
			Duration: constant.DayDuration(),
		},
		WebsiteUrl: convertedWebsiteUrl,
		Coordinate: value.GeoPoint{
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		},
		Images:   imageModels,
		Displays: displayModels,
	}
}
