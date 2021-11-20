package model

import "harvest/src/domain/value"

type SpaceDisplayer struct {
	Id          value.SpaceDisplayerId
	ImageUrl    string
	Description string
}
