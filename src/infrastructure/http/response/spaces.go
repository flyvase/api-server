package response

import (
	"api-server/src/domain/model"
	"encoding/json"
)

type space struct {
	Id               uint32            `json:"id"`
	Headline         string            `json:"headline"`
	Access           string            `json:"access,omitempty"`
	NumberOfVisitors *numberOfVisitors `json:"number_of_visitors,omitempty"`
	CustomerSegment  *customerSegment  `json:"customer_segment,omitempty"`
	Price            price             `json:"price"`
	Coordinate       geoPoint          `json:"coordinate"`
	Images           []*spaceImage     `json:"images"`
}

func spaceFromModel(s *model.Space) *space {
	var images []*spaceImage
	for _, i := range s.Images {
		images = append(images, spaceImageFromModel(i))
	}

	if images == nil {
		images = []*spaceImage{}
	}

	numberOfVisitors := numberOfVisitorsFromValue(s.NumberOfVisitors)
	if s.NumberOfVisitors.IsEmpty() {
		numberOfVisitors = nil
	}

	customerSegment := customerSegmentFromValue(s.CustomerSegment)
	if s.CustomerSegment.IsEmpty() {
		customerSegment = nil
	}

	return &space{
		Id:               uint32(s.Id.Value),
		Headline:         s.Headline,
		Access:           s.Access,
		NumberOfVisitors: numberOfVisitors,
		CustomerSegment:  customerSegment,
		Price: priceFromValue(
			s.Price,
		),
		Coordinate: geoPointFromValue(
			s.Coordinate,
		),
		Images: images,
	}
}

type spaces struct {
	List []*space `json:"list"`
}

func EncodeSpaceModels(spaceModels []*model.Space) ([]byte, error) {
	var list []*space
	for _, sm := range spaceModels {
		list = append(list, spaceFromModel(sm))
	}

	if list == nil {
		list = []*space{}
	}

	spaces := &spaces{
		List: list,
	}

	return json.Marshal(spaces)
}
