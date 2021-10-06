package response

import (
	"encoding/json"
	"harvest/domain/entity"
	"net/http"
)

type Space struct {
	Id   uint32 `json:"id" validate:"required,max=255"`
	Name string `json:"name" validate:"required,max=150"`
}

func MarshalSpaceResponseJson(w http.ResponseWriter, se []entity.Space) error {
	var spaces []Space
	for _, s := range se {
	  var ms Space = Space{s.Id, s.Name}
		spaces = append(spaces, ms)
	}

	js, err := json.Marshal(spaces)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
