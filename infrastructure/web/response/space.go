package response

import (
	"encoding/json"
	"harvest/domain/entity"
	"net/http"
)

type Space struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func MarshalSpaceResponseJson(w http.ResponseWriter, entities []entity.Space) ([]byte, error) {
	var spaces []Space
	for _, se := range entities {
	  s := Space{se.Id, se.Name}
		spaces = append(spaces, s)
	}

	js, err := json.Marshal(spaces)

	return js, err
}
