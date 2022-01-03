package entity

import (
	"api-server/src/domain/model"
	"api-server/src/domain/value"
)

type SpaceImage struct {
	Id       uint64
	ImageUrl string
}

func (i *SpaceImage) toSpaceImageModel() *model.SpaceImage {
	return &model.SpaceImage{
		Id: value.SpaceImageId{
			Value: i.Id,
		},
		ImageUrl: i.ImageUrl,
	}
}
