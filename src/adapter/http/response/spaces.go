package response

import (
	"encoding/json"
	"harvest/src/domain/model"
)

type space struct {
	Id               uint32            `json:"id"`
	Headline         string            `json:"headline"`
	Access           string            `json:"access"`
	NumberOfVisitors *numberOfVisitors `json:"number_of_visitors"`
	CustomerSegment  *customerSegment  `json:"customer_segment"`
	Price            *price            `json:"price"`
	Images           []*spaceImage     `json:"images"`
}

func spaceFromModel(s *model.Space) *space {
	var images []*spaceImage
	for _, i := range s.Images {
		images = append(images, spaceImageFromModel(i))
	}

	return &space{
		Id:       uint32(s.Id.Value),
		Headline: s.Headline,
		Access:   s.Access,
		NumberOfVisitors: numberOfVisitorsFromValue(
			&s.NumberOfVisitors,
		),
		CustomerSegment: customerSegmentFromValue(
			&s.CustomerSegment,
		),
		Price: priceFromValue(
			&s.Price,
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

	spaces := &spaces{
		List: list,
	}

	return json.Marshal(spaces)
}
