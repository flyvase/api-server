package space

import "harvest/src/domain/value/core"

type CustomerSegment struct {
	Sex    core.Sex
	MinAge uint8
	MaxAge uint8
}
