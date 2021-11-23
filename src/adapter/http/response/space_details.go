package response

import (
	"encoding/json"
	"harvest/src/domain/model"
)

type spaceDetails struct {
	Id               uint32            `json:"id"`
	Headline         string            `json:"headline"`
	Access           string            `json:"access"`
	NumberOfVisitors numberOfVisitors  `json:"number_of_visitors"`
	CustomerSegment  customerSegment   `json:"customer_segment"`
	Price            price             `json:"price"`
	WebsiteUrl       string            `json:"website_url"`
	Coordinate       geoPoint          `json:"coordinate"`
	Images           []*spaceImage     `json:"images"`
	Displayers       []*spaceDisplayer `json:"displayers"`
}

func spaceDetailsFromModel(s *model.Space) *spaceDetails {
	var images []*spaceImage
	for _, i := range s.Images {
		images = append(images, spaceImageFromModel(i))
	}

	if images == nil {
		images = []*spaceImage{}
	}

	var displayers []*spaceDisplayer
	for _, d := range s.Displayers {
		displayers = append(displayers, spaceDisplayerFromModel(d))
	}

	if displayers == nil {
		displayers = []*spaceDisplayer{}
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
		Images:     images,
		Displayers: displayers,
	}
}

func EncodeSpaceModel(spaceModel *model.Space) ([]byte, error) {
	details := spaceDetailsFromModel(spaceModel)

	return json.Marshal(details)
}
