package response

import "harvest/src/domain/model"

type spaceDisplayer struct {
	Id          uint32 `json:"id"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

func spaceDisplayerFromModel(d *model.SpaceDisplayer) *spaceDisplayer {
	return &spaceDisplayer{
		Id:          uint32(d.Id.Value),
		ImageUrl:    d.ImageUrl,
		Description: d.Description,
	}
}
