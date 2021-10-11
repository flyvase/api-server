package response

import (
	"encoding/json"
	"harvest/domain/entity"
)

type Space struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func EncodeSpaceEntities(entities []entity.Space) ([]byte, error) {
	var spaces []Space
	for _, se := range entities {
	    s := Space{se.Id, se.Name}
		spaces = append(spaces, s)
	}

	js, err := json.Marshal(spaces)

	if err != nil {
  	    return nil, err
	}

	return js, nil
}
