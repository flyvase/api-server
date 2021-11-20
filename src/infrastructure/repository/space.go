package repository

import (
	"encoding/json"
	"harvest/src/core/constants"
	"harvest/src/domain/model"
	"harvest/src/domain/value/core"
	"harvest/src/domain/value/space"
	"harvest/src/domain/value/spaceimage"
	"harvest/src/infrastructure/sql"
	"harvest/src/infrastructure/sql/entity"
)

type image struct {
	Id       uint32 `json:"id"`
	ImageUrl string `json:"image_url"`
}

func (i *image) toSpaceImageModel() *model.SpaceImage {
	return &model.SpaceImage{
		Id: spaceimage.Id{
			Value: i.Id,
		},
		ImageUrl: i.ImageUrl,
	}
}

type listResult struct {
	entity.Space
	ImagesJson []byte
}

func (r *listResult) getImages() []*image {
	var images []*image
	json.Unmarshal(r.ImagesJson, &images)
	return images
}

func (r *listResult) toSpaceModel() *model.Space {
	images := r.getImages()

	var imageModels []*model.SpaceImage
	for _, i := range images {
		imageModels = append(imageModels, i.toSpaceImageModel())
	}

	return &model.Space{
		Id: space.Id{
			Value: r.Id,
		},
		Headline: r.Headline,
		Access: core.Access{
			Value: r.Access,
		},
		NumberOfVisitors: space.NumberOfVisitors{
			Visitors: uint(r.WeeklyVisitors),
			Duration: constants.WeekDuration(),
		},
		CustomerSegment: space.CustomerSegment{
			Sex:    core.SexFromString(r.MainCustomersSex),
			MinAge: r.MinMainCustomersAge,
			MaxAge: r.MaxMainCustomersAge,
		},
		Price: space.Price{
			Value:    uint(r.DailyPrice),
			Duration: constants.DayDuration(),
		},
		Images: imageModels,
	}
}

type Space struct {
	SqlDriver sql.Driver
}

func (s *Space) List() ([]*model.Space, error) {
	spacesRows, err := s.SqlDriver.Query(`
		select
		spaces.id,
		spaces.headline,
		spaces.access,
		spaces.weekly_visitors,
		spaces.main_customers_sex,
		spaces.min_main_customers_age,
		spaces.max_main_customers_age,
		spaces.daily_price,
		json_arrayagg(json_object("id", space_images.id, "image_url", space_images.image_url))
		from spaces
		inner join space_images on spaces.id = space_images.space_id
		group by spaces.id
	`)
	if err != nil {
		return nil, err
	}

	defer spacesRows.Close()

	var results []*listResult
	for spacesRows.Next() {
		var result listResult
		spacesRows.Scan(
			&result.Id,
			&result.Headline,
			&result.Access,
			&result.WeeklyVisitors,
			&result.MainCustomersSex,
			&result.MinMainCustomersAge,
			&result.MaxMainCustomersAge,
			&result.DailyPrice,
			&result.ImagesJson,
		)

		results = append(results, &result)
	}

	var spaceModels []*model.Space
	for _, s := range results {
		spaceModels = append(spaceModels, s.toSpaceModel())
	}

	return spaceModels, nil
}
