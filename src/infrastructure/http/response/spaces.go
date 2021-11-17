package response

import (
	"encoding/json"
	"harvest/src/domain/entity"
	"harvest/src/domain/object"
)

type Space struct {
	Id                  uint32 `json:"id"`
	Headline            string `json:"headline"`
	Access              string `json:"access"`
	NumberOfVisitors    uint32 `json:"number_of_visitors"`
	MainCustomersSex    string `json:"main_customers_sex"` // using string temporarily
	MinMainCustomersAge uint8  `json:"min_main_customers_age"`
	MaxMainCustomersAge uint8  `json:"max_main_customers_age"`
	Price               uint32 `json:"price"`
}

type Spaces struct {
	List []*Space `json:"list"`
}

type SpaceDetail struct {
	Id                  uint32           `json:"id"`
	Headline            string           `json:"headline"`
	Access              string           `json:"access"`
	NumberOfVisitors    uint32           `json:"number_of_visitors"`
	MainCustomersSex    string           `json:"main_customers_sex"` // using string temporarily
	MinMainCustomersAge uint8            `json:"min_main_customers_age"`
	MaxMainCustomersAge uint8            `json:"max_main_customers_age"`
	Price               uint32           `json:"price"`
	WebsiteUrl          string           `json:"website_url"`
	Coordinate          object.GeoPoint  `json:"coordinate"`
	Images              []SpaceImage     `json:"images"`
	Displayers          []SpaceDisplayer `json:"displayers"`
}

type SpaceImage struct {
	Id       uint32 `json:"id"`
	ImageUrl string `json:"image_url"`
}

type SpaceDisplayer struct {
	Id          uint32 `json:"id"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
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
	spaceImages := EncodeSpaceImageEntities(entity.Images)
	spaceDisplayer := EncodeSpaceDisplayerEntities(entity.Displayers)
	s := SpaceDetail{entity.Id, entity.Headline, entity.Access, entity.NumberOfVisitors, entity.MainCustomersSex, entity.MinMainCustomersAge, entity.MaxMainCustomersAge, entity.Price, entity.WebsiteUrl, entity.Coordinate, spaceImages, spaceDisplayer}
	return json.Marshal(s)
}

func EncodeSpaceImageEntities(entities []*entity.SpaceImage) []SpaceImage {
	var spaceImages []SpaceImage
	for _, sie := range entities {
		si := SpaceImage{sie.Id, sie.ImageUrl}
		spaceImages = append(spaceImages, si)
	}

	return spaceImages
}

func EncodeSpaceDisplayerEntities(entities []*entity.SpaceDisplayer) []SpaceDisplayer {
	var SpaceDisplayers []SpaceDisplayer
	for _, sde := range entities {
		sd := SpaceDisplayer{sde.Id, sde.Description, sde.ImageUrl}
		SpaceDisplayers = append(SpaceDisplayers, sd)
	}

	return SpaceDisplayers
}
