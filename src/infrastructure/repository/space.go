package repository

import (
	"encoding/json"
	"harvest/src/core/constants"
	"harvest/src/domain/model"
	"harvest/src/domain/value"
	"harvest/src/infrastructure/entity"
	"harvest/src/infrastructure/sql"
	"strconv"
)

type image struct {
	Id       uint32 `json:"id"`
	ImageUrl string `json:"image_url"`
}

func (i *image) toSpaceImageModel() *model.SpaceImage {
	return &model.SpaceImage{
		Id: value.SpaceImageId{
			Value: uint(i.Id),
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

	c, err := strconv.Atoi(r.MainCustomersSex)
	if err != nil {
		panic(err)
	}
	sexCode := uint8(c)

	return &model.Space{
		Id: value.SpaceId{
			Value: uint(r.Id),
		},
		Headline: r.Headline,
		Access:   r.Access,
		NumberOfVisitors: value.NumberOfVisitors{
			Visitors: uint(r.WeeklyVisitors),
			Duration: constants.WeekDuration(),
		},
		CustomerSegment: value.CustomerSegment{
			Sex:    value.NewSex(sexCode),
			MinAge: r.MinMainCustomersAge,
			MaxAge: r.MaxMainCustomersAge,
		},
		Price: value.Price{
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
		if err := spacesRows.Scan(
			&result.Id,
			&result.Headline,
			&result.Access,
			&result.WeeklyVisitors,
			&result.MainCustomersSex,
			&result.MinMainCustomersAge,
			&result.MaxMainCustomersAge,
			&result.DailyPrice,
			&result.ImagesJson,
		); err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	var spaceModels []*model.Space
	for _, s := range results {
		spaceModels = append(spaceModels, s.toSpaceModel())
	}

	return spaceModels, nil
}
