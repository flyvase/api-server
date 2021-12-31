package response

import "api-server/src/domain/value"

type customerSegment struct {
	Sex    string `json:"sex"`
	MinAge uint8  `json:"min_age"`
	MaxAge uint8  `json:"max_age"`
}

func customerSegmentFromValue(s value.CustomerSegment) customerSegment {
	return customerSegment{
		Sex:    s.Sex.ToString(),
		MinAge: s.MinAge,
		MaxAge: s.MaxAge,
	}
}
