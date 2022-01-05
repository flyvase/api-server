package response

import (
	"api-server/src/domain/model"
	"encoding/json"
)

type spaceDetails struct {
	Id               uint32            `json:"id"`
	Headline         string            `json:"headline"`
	Access           string            `json:"access,omitempty"`
	NumberOfVisitors *numberOfVisitors `json:"number_of_visitors,omitempty"`
	CustomerSegment  *customerSegment  `json:"customer_segment,omitempty"`
	Price            price             `json:"price"`
	WebsiteUrl       string            `json:"website_url,omitempty"`
	Coordinate       geoPoint          `json:"coordinate"`
	Images           []*spaceImage     `json:"images"`
	Displays         []*spaceDisplay   `json:"displays"`
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

	numberOfVisitors := numberOfVisitorsFromValue(s.NumberOfVisitors)
	if s.NumberOfVisitors.IsEmpty() {
		numberOfVisitors = nil
	}

	customerSegment := customerSegmentFromValue(s.CustomerSegment)
	if s.CustomerSegment.IsEmpty() {
		customerSegment = nil
	}

	return &spaceDetails{
		Id:               uint32(s.Id.Value),
		Headline:         s.Headline,
		Access:           s.Access,
		NumberOfVisitors: numberOfVisitors,
		CustomerSegment:  customerSegment,
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
