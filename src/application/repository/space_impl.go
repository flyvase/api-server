package repository

import (
	"encoding/json"
	"harvest/src/application/gateway/entity"
	"harvest/src/application/gateway/sql"
	"harvest/src/core/constant"
	"harvest/src/core/errors"
	"harvest/src/domain/model"
	"harvest/src/domain/value"
	"strconv"
	"sync"
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

func isEmptyImages(images []*image) bool {
	if len(images) == 1 && images[0].Id == 0 && images[0].ImageUrl == "" {
		return true
	}

	return false
}

func (r *listResult) decodeImages() []*image {
	var images []*image
	json.Unmarshal(r.ImagesJson, &images)

	if isEmptyImages(images) {
		return nil
	}

	return images
}

func (r *listResult) toSpaceModel() *model.Space {
	images := r.decodeImages()

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
			Duration: constant.WeekDuration(),
		},
		CustomerSegment: value.CustomerSegment{
			Sex:    value.NewSex(sexCode),
			MinAge: r.MinMainCustomersAge,
			MaxAge: r.MaxMainCustomersAge,
		},
		Price: value.Price{
			Price:    uint(r.DailyPrice),
			Duration: constant.DayDuration(),
		},
		Images: imageModels,
	}
}

type SpaceImpl struct {
	SqlDriver sql.Driver
}

func (s *SpaceImpl) List() ([]*model.Space, error) {
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
		left join space_images on spaces.id = space_images.space_id
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

func (s *SpaceImpl) getSpace(id value.SpaceId) (*entity.Space, error) {
	spaceRow := s.SqlDriver.QueryRow(`
		select
		id,
		headline,
		access,
		weekly_visitors,
		main_customers_sex,
		min_main_customers_age,
		max_main_customers_age,
		daily_price,
		website_url,
		ST_Latitude(coordinate),
		ST_Longitude(coordinate)
		from spaces
		where id = ?
	`, id.Value)

	var spaceEntity entity.Space
	if err := spaceRow.Scan(
		&spaceEntity.Id,
		&spaceEntity.Headline,
		&spaceEntity.Access,
		&spaceEntity.WeeklyVisitors,
		&spaceEntity.MainCustomersSex,
		&spaceEntity.MinMainCustomersAge,
		&spaceEntity.MaxMainCustomersAge,
		&spaceEntity.DailyPrice,
		&spaceEntity.WebsiteUrl,
		&spaceEntity.Latitude,
		&spaceEntity.Longitude,
	); err != nil {
		return nil, err
	}

	return &spaceEntity, nil
}

type getSpaceImagesResult struct {
	Value []*entity.SpaceImage
	Error error
}

func (s *SpaceImpl) getSpaceImages(id value.SpaceId, c chan *getSpaceImagesResult, wg *sync.WaitGroup) {
	defer wg.Done()

	spaceImagesRows, err := s.SqlDriver.Query(`
		select
		id,
		image_url
		from space_images
		where space_id = ?
	`, id.Value)
	if err != nil {
		c <- &getSpaceImagesResult{
			Value: nil,
			Error: err,
		}
	}

	defer spaceImagesRows.Close()

	var spaceImageEntities []*entity.SpaceImage
	for spaceImagesRows.Next() {
		var spaceImageEntity entity.SpaceImage
		if err := spaceImagesRows.Scan(
			&spaceImageEntity.Id,
			&spaceImageEntity.ImageUrl,
		); err != nil {
			c <- &getSpaceImagesResult{
				Value: nil,
				Error: err,
			}
		}

		spaceImageEntities = append(spaceImageEntities, &spaceImageEntity)
	}

	c <- &getSpaceImagesResult{
		Value: spaceImageEntities,
		Error: nil,
	}
}

type getSpaceDisplayersResult struct {
	Value []*entity.SpaceDisplayer
	Error error
}

func (s *SpaceImpl) getSpaceDisplayers(id value.SpaceId, c chan *getSpaceDisplayersResult, wg *sync.WaitGroup) {
	defer wg.Done()

	spaceDisplayersRows, err := s.SqlDriver.Query(`
		select
		id,
		image_url,
		description
		from space_displayers
		where space_id = ?
	`, id.Value)
	if err != nil {
		c <- &getSpaceDisplayersResult{
			Value: nil,
			Error: err,
		}
	}

	defer spaceDisplayersRows.Close()

	var spaceDisplayerEntities []*entity.SpaceDisplayer
	for spaceDisplayersRows.Next() {
		var spaceDisplayerEntity entity.SpaceDisplayer
		if err := spaceDisplayersRows.Scan(
			&spaceDisplayerEntity.Id,
			&spaceDisplayerEntity.ImageUrl,
			&spaceDisplayerEntity.Description,
		); err != nil {
			c <- &getSpaceDisplayersResult{
				Value: nil,
				Error: err,
			}
		}

		spaceDisplayerEntities = append(spaceDisplayerEntities, &spaceDisplayerEntity)
	}

	c <- &getSpaceDisplayersResult{
		Value: spaceDisplayerEntities,
		Error: nil,
	}
}

func (s *SpaceImpl) Fetch(id value.SpaceId) (*model.Space, error) {
	spaceEntity, err := s.getSpace(id)
	if err != nil {
		if err == errors.ErrSqlNoRows {
			return nil, nil
		}

		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	spaceImagesChannel := make(chan *getSpaceImagesResult, 1)
	spaceDisplayersChannel := make(chan *getSpaceDisplayersResult, 1)

	go s.getSpaceImages(id, spaceImagesChannel, &wg)
	go s.getSpaceDisplayers(id, spaceDisplayersChannel, &wg)

	wg.Wait()

	getSpaceImagesResult := <-spaceImagesChannel
	if getSpaceImagesResult.Error != nil {
		return nil, getSpaceImagesResult.Error
	}

	getSpaceDisplayersResult := <-spaceDisplayersChannel
	if getSpaceDisplayersResult.Error != nil {
		return nil, getSpaceDisplayersResult.Error
	}

	return spaceEntity.ToSpaceModel(
		getSpaceImagesResult.Value,
		getSpaceDisplayersResult.Value,
	), nil
}
