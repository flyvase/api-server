package value

import "time"

type NumberOfVisitors struct {
	Visitors uint
	Duration time.Duration
}

func (n *NumberOfVisitors) IsEmpty() bool {
	if n.Visitors == 0 && n.Duration.Milliseconds() == 0 {
		return true
	}

	return false
}
