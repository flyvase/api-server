package repository

import (
	"api-server/src/core/constant"
	"api-server/src/domain/model"
	"api-server/src/domain/value"
	"api-server/src/infrastructure/entity"
	"api-server/src/infrastructure/gateway/sql"
	stdsql "database/sql"
	"encoding/json"
	"strconv"
	"sync"
)

type listResult struct {
	entity.Space
	ImagesJson []byte
}

func isEmptyImages(images []*entity.SpaceImage) bool {
	if len(images) == 1 && images[0].Id == 0 && images[0].ImageUrl == "" {
		return true
	}

	return false
}

func (r *listResult) decodeImages() []*entity.SpaceImage {
	var images []*entity.SpaceImage
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
		imageModels = append(imageModels, i.ToSpaceImageModel())
	}

	var access string
	if r.Access.Valid {
		access = r.Access.String
	}

	var numberOfVisitors value.NumberOfVisitors
	if r.WeeklyVisitors.Valid {
		numberOfVisitors = value.NumberOfVisitors{
			Visitors: uint32(r.WeeklyVisitors.Int32),
			Duration: constant.WeekDuration(),
		}
	}

	c, err := strconv.Atoi(r.MainCustomersSex)
	if err != nil {
		panic(err)
	}
	sexCode := uint8(c)

	var customerSegment value.CustomerSegment
	if r.MinMainCustomersAge.Valid && r.MaxMainCustomersAge.Valid {
		customerSegment = value.CustomerSegment{
			Sex: value.Sex{
				Code: sexCode,
			},
			MinAge: uint8(r.MinMainCustomersAge.Int32),
			MaxAge: uint8(r.MaxMainCustomersAge.Int32),
		}
	}

	return &model.Space{
		Id: value.SpaceId{
			Value: r.Id,
		},
		Headline:         r.Headline,
		Access:           access,
		NumberOfVisitors: numberOfVisitors,
		CustomerSegment:  customerSegment,
		Price: value.Price{
			Price:    r.DailyPrice,
			Duration: constant.DayDuration(),
		},
		Coordinate: value.GeoPoint{
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
		},
		Images: imageModels,
	}
}

type SpaceImpl struct {
	SqlDriver *sql.Driver
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
		ST_Latitude(coordinate),
		ST_Longitude(coordinate),
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
			&result.Latitude,
			&result.Longitude,
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

type getSpaceDisplaysResult struct {
	Value []*entity.SpaceDisplay
	Error error
}

func (s *SpaceImpl) getSpaceDisplays(id value.SpaceId, c chan *getSpaceDisplaysResult, wg *sync.WaitGroup) {
	defer wg.Done()

	spaceDisplaysRows, err := s.SqlDriver.Query(`
		select
		id,
		image_url,
		description
		from space_displays
		where space_id = ?
	`, id.Value)
	if err != nil {
		c <- &getSpaceDisplaysResult{
			Value: nil,
			Error: err,
		}
	}

	defer spaceDisplaysRows.Close()

	var spaceDisplayEntities []*entity.SpaceDisplay
	for spaceDisplaysRows.Next() {
		var spaceDisplayEntity entity.SpaceDisplay
		if err := spaceDisplaysRows.Scan(
			&spaceDisplayEntity.Id,
			&spaceDisplayEntity.ImageUrl,
			&spaceDisplayEntity.Description,
		); err != nil {
			c <- &getSpaceDisplaysResult{
				Value: nil,
				Error: err,
			}
		}

		spaceDisplayEntities = append(spaceDisplayEntities, &spaceDisplayEntity)
	}

	c <- &getSpaceDisplaysResult{
		Value: spaceDisplayEntities,
		Error: nil,
	}
}

func (s *SpaceImpl) Fetch(id value.SpaceId) (*model.Space, error) {
	spaceEntity, err := s.getSpace(id)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	spaceImagesChannel := make(chan *getSpaceImagesResult, 1)
	spaceDisplaysChannel := make(chan *getSpaceDisplaysResult, 1)

	go s.getSpaceImages(id, spaceImagesChannel, &wg)
	go s.getSpaceDisplays(id, spaceDisplaysChannel, &wg)

	wg.Wait()

	getSpaceImagesResult := <-spaceImagesChannel
	if getSpaceImagesResult.Error != nil {
		return nil, getSpaceImagesResult.Error
	}

	getSpaceDisplaysResult := <-spaceDisplaysChannel
	if getSpaceDisplaysResult.Error != nil {
		return nil, getSpaceDisplaysResult.Error
	}

	return spaceEntity.ToSpaceModel(
		getSpaceImagesResult.Value,
		getSpaceDisplaysResult.Value,
	), nil
}

func (s *SpaceImpl) GetWebsiteUrl(id value.SpaceId) (string, error) {
	row := s.SqlDriver.QueryRow(`
		select
		website_url
		from spaces
		where id = ?
	`, id.Value)

	var url stdsql.NullString
	if err := row.Scan(
		&url,
	); err != nil {
		return "", err
	}

	return url.String, nil
}
