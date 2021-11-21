package response

import "harvest/src/domain/model"

type spaceImage struct {
	Id       uint32 `json:"id"`
	ImageUrl string `json:"image_url"`
}

func spaceImageFromModel(s *model.SpaceImage) *spaceImage {
	return &spaceImage{
		Id:       uint32(s.Id.Value),
		ImageUrl: s.ImageUrl,
	}
}
