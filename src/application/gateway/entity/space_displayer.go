package entity

import (
	"harvest/src/domain/model"
	"harvest/src/domain/value"
)

type SpaceDisplayer struct {
	Id          uint32
	ImageUrl    string
	Description string
}

func (d *SpaceDisplayer) toSpaceDisplayerModel() *model.SpaceDisplayer {
	return &model.SpaceDisplayer{
		Id: value.SpaceDisplayerId{
			Value: uint(d.Id),
		},
		ImageUrl:    d.ImageUrl,
		Description: d.Description,
	}
}
