package response

import (
	"encoding/json"
	"harvest/src/domain/model"
)

type spaceDetails struct {
	Id               uint32           `json:"id"`
	Headline         string           `json:"headline"`
	Access           string           `json:"access"`
	NumberOfVisitors numberOfVisitors `json:"number_of_visitors"`
	CustomerSegment  customerSegment  `json:"customer_segment"`
	Price            price            `json:"price"`
	WebsiteUrl       string           `json:"website_url"`
	Coordinate       geoPoint         `json:"coordinate"`
	Images           []*spaceImage    `json:"images"`
	Displays         []*spaceDisplay  `json:"displays"`
}

func spaceDetailsFromModel(s *model.Space) *spaceDetails {
	var images []*spaceImage
	for _, i := range s.Images {
		images = append(images, spaceImageFromModel(i))
	}

	if images == nil {
		images = []*spaceImage{}
	}

	var displays []*spaceDisplay
	for _, d := range s.Displays {
		displays = append(displays, spaceDisplayFromModel(d))
	}

	if displays == nil {
		displays = []*spaceDisplay{}
	}

	return &spaceDetails{
		Id:       uint32(s.Id.Value),
		Headline: s.Headline,
		Access:   s.Access,
		NumberOfVisitors: numberOfVisitorsFromValue(
			s.NumberOfVisitors,
		),
		CustomerSegment: customerSegmentFromValue(
			s.CustomerSegment,
		),
		Price: priceFromValue(
			s.Price,
		),
		WebsiteUrl: s.WebsiteUrl,
		Coordinate: geoPoint{
			Latitude:  s.Coordinate.Latitude,
			Longitude: s.Coordinate.Longitude,
		},
		Images:   images,
		Displays: displays,
	}
}

func EncodeSpaceModel(spaceModel *model.Space) ([]byte, error) {
	details := spaceDetailsFromModel(spaceModel)

	return json.Marshal(details)
}
