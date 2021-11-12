package repository

import (
	"harvest/src/domain/entity"
	"harvest/src/infrastructure/sql"
)

type Space struct {
	Driver sql.Driver
}

func (sr *Space) List() ([]*entity.Space, error) {
	rows, err := sr.Driver.Query(
		"select id, headline, access, number_of_visitors, main_customers_sex, min_main_customers_age, max_main_customers_age, price from spaces",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var spaces []*entity.Space

	for rows.Next() {
		var se entity.Space
		if err := rows.Scan(&se.Id, &se.Headline, &se.Access, &se.NumberOfVisitors, &se.MainCustomersSex, &se.MinMainCustomersAge, &se.MaxMainCustomersAge, &se.Price); err != nil {
			return nil, err
		}

		spaces = append(spaces, &se)
	}

	return spaces, nil
}

func (sr *Space) Fetch(id uint32) (*entity.Space, error) {
	spacesRow, err := sr.Driver.QueryRow(
		`select id, headline, access, number_of_visitors, main_customers_sex, min_main_customers_age, max_main_customers_age, price, website_url from spaces where id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var se entity.Space
	if err := spacesRow.Scan(&se.Id, &se.Headline, &se.Access, &se.NumberOfVisitors, &se.MainCustomersSex, &se.MinMainCustomersAge, &se.MaxMainCustomersAge, &se.Price, &se.WebsiteUrl); err != nil {
		return nil, err
	}

	spaceImagesRows, err := sr.Driver.Query(
		`select id, image_url from space_images where space_id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer spaceImagesRows.Close()

	var spaceImages []*entity.SpaceImage
	for spaceImagesRows.Next() {
		var sie entity.SpaceImage
		if err := spaceImagesRows.Scan(&sie.Id, &sie.ImageUrl); err != nil {
			return nil, err
		}
		spaceImages = append(spaceImages, &sie)
	}
	se.Images = spaceImages

	spaceDisplayersRows, err := sr.Driver.Query(
		`select id, image_url, description from space_displayers where space_id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer spaceDisplayersRows.Close()

	var spaceDisplayers []*entity.SpaceDisplayer
	for spaceDisplayersRows.Next() {
		var sde entity.SpaceDisplayer
		if err := spaceDisplayersRows.Scan(&sde.Id, &sde.ImageUrl, &sde.Description); err != nil {
			return nil, err
		}
		spaceDisplayers = append(spaceDisplayers, &sde)
	}
	se.Displayers = spaceDisplayers

	return &se, nil
}
