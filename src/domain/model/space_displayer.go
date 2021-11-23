package model

import "harvest/src/domain/value"

type SpaceDisplay struct {
	Id          value.SpaceDisplayId
	ImageUrl    string
	Description string
}
