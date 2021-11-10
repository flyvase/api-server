package response

import (
	"encoding/json"
	"harvest/src/domain/entity"
	"harvest/src/domain/object"
)

type Space struct {
	Id                  uint32 `json:"id"`
	Headline            string `json:"headline"`
	Access              string
	NumberOfVisitors    uint32
	MainCustomersSex    string // using string temporarily
	MinMainCustomersAge uint8
	MaxMainCustomersAge uint8
	Price               uint32
}

type Spaces struct {
	List []*Space `json:"list"`
}

type SpaceDetail struct {
	Id                  uint32
	Headline            string
	Access              string
	NumberOfVisitors    uint32
	MainCustomersSex    string // using string temporarily
	MinMainCustomersAge uint8
	MaxMainCustomersAge uint8
	Price               uint32
	WebsiteUrl          string
	Coordinate          object.GeoPoint
	Images              []*entity.SpaceImage
	Displayers          []*entity.SpaceDisplayer
}

func EncodeSpaceEntities(entities []*entity.Space) ([]byte, error) {
	var list []*Space
	for _, se := range entities {
		s := Space{se.Id, se.Headline, se.Access, se.NumberOfVisitors, se.MainCustomersSex, se.MinMainCustomersAge, se.MaxMainCustomersAge, se.Price}
		list = append(list, &s)
	}

	spaces := Spaces{List: list}

	js, err := json.Marshal(spaces)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func EncodeSpaceEntity(entity *entity.Space) ([]byte, error) {
	s := SpaceDetail{entity.Id, entity.Headline, entity.Access, entity.NumberOfVisitors, entity.MainCustomersSex, entity.MinMainCustomersAge, entity.MaxMainCustomersAge, entity.Price, entity.WebsiteUrl, entity.Coordinate, entity.Images, entity.Displayers}
	return json.Marshal(s)
}
