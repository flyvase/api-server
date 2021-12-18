package entity

import (
	"harvest/src/domain/model"
	"harvest/src/domain/value"
)

type SpaceImage struct {
	Id       uint32
	ImageUrl string
}

func (i *SpaceImage) toSpaceImageModel() *model.SpaceImage {
	return &model.SpaceImage{
		Id: value.SpaceImageId{
			Value: uint(i.Id),
		},
		ImageUrl: i.ImageUrl,
	}
}