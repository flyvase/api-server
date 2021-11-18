package model

import "harvest/src/domain/value/spacedisplayer"

type SpaceDisplayer struct {
	Id          spacedisplayer.Id
	ImageUrl    string
	Description string
}
