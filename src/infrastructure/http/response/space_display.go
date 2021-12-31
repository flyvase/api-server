package response

import "api-server/src/domain/model"

type spaceDisplay struct {
	Id          uint32 `json:"id"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

func spaceDisplayFromModel(d *model.SpaceDisplay) *spaceDisplay {
	return &spaceDisplay{
		Id:          uint32(d.Id.Value),
		ImageUrl:    d.ImageUrl,
		Description: d.Description,
	}
}
