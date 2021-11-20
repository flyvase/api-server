package main

import (
	"fmt"
	"harvest/src/infrastructure/repository"
	"harvest/src/infrastructure/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlDriver := sql.NewDriver()

	repo := repository.Space{
		SqlDriver: sqlDriver,
	}

	res, _ := repo.List()
	for _, v := range res {
		fmt.Println(v.Id)
		fmt.Println(v.Headline)
		fmt.Println(v.Access)
		fmt.Println(v.NumberOfVisitors)
		fmt.Println(v.CustomerSegment)
		fmt.Println(v.Price)
		for _, i := range v.Images {
			fmt.Println(i)
		}
	}
}

// func (s *Space) Fetch(id space.Id) (*model.Space, error) {
// 	spaceRow := s.SqlDriver.QueryRow(`
// 		select
// 		id,
// 		headline,
// 		access,
// 		weekly_visitors,
// 		main_customers_sex,
// 		min_main_customers_age,
// 		max_main_customers_age,
// 		daily_price,
// 		website_url,
// 		ST_Latitude(coordinate),
// 		ST_Longitude(coordinate)
// 		from spaces
// 		where id = ?
// 	`, id.Value)

// 	var spaceEntity *entity.Space
// 	if err := spaceRow.Scan(
// 		spaceEntity.Id,
// 		spaceEntity.Headline,
// 		spaceEntity.Access,
// 		spaceEntity.WeeklyVisitors,
// 		spaceEntity.MainCustomersSex,
// 		spaceEntity.MinMainCustomersAge,
// 		spaceEntity.MaxMainCustomersAge,
// 		spaceEntity.DailyPrice,
// 		spaceEntity.WebsiteUrl,
// 		spaceEntity.Latitude,
// 		spaceEntity.Longitude,
// 	); err != nil {
// 		return nil, err
// 	}

// 	spaceImagesRows, err := s.SqlDriver.Query(`
// 		select
// 		id,
// 		image_url
// 		from space_images
// 		where space_id = ?
// 	`, id.Value)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer spaceImagesRows.Close()

// 	var spaceImageEntities []*entity.SpaceImage
// 	for spaceImagesRows.Next() {
// 		var spaceImageEntity entity.SpaceImage
// 		spaceImagesRows.Scan(
// 			&spaceImageEntity.Id,
// 			&spaceImageEntity.ImageUrl,
// 		)

// 		spaceImageEntities = append(spaceImageEntities, &spaceImageEntity)
// 	}

// 	spaceDisplayersRows, err := s.SqlDriver.Query(`
// 		select
// 		id,
// 		image_url,
// 		description
// 		from space_displayers
// 		where space_id = ?
// 	`, id.Value)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var spaceDisplayerEntities []*entity
// }
