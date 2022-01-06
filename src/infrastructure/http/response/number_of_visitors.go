package response

import "api-server/src/domain/value"

type numberOfVisitors struct {
	Visitors uint32 `json:"visitors"`
	Duration int64  `json:"duration"`
}

func numberOfVisitorsFromValue(n value.NumberOfVisitors) *numberOfVisitors {
	return &numberOfVisitors{
		Visitors: n.Visitors,
		Duration: n.Duration.Milliseconds(),
	}
}
