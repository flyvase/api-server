package entity

import (
	"api-server/src/domain/model"
	"api-server/src/domain/value"
)

type SpaceDisplay struct {
	Id          uint64
	ImageUrl    string
	Description string
}

func (d *SpaceDisplay) toSpaceDisplayModel() *model.SpaceDisplay {
	return &model.SpaceDisplay{
		Id: value.SpaceDisplayId{
			Value: d.Id,
		},
		ImageUrl:    d.ImageUrl,
		Description: d.Description,
	}
}
