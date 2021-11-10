package response

import (
	"encoding/json"
	"harvest/src/domain/entity"
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
	s := Space{entity.Id, entity.Headline}
	return json.Marshal(s)
}
