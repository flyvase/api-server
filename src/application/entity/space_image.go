package entity

import (
	"api-server/src/domain/model"
	"api-server/src/domain/value"
)

type SpaceImage struct {
	Id       uint64 `json:"id"`
	ImageUrl string `json:"image_url"`
}

func (i *SpaceImage) ToSpaceImageModel() *model.SpaceImage {
	return &model.SpaceImage{
		Id: value.SpaceImageId{
			Value: i.Id,
		},
		ImageUrl: i.ImageUrl,
	}
}
