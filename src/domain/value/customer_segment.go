package value

type CustomerSegment struct {
	Sex    Sex
	MinAge uint8
	MaxAge uint8
}

func (c *CustomerSegment) IsEmpty() bool {
	if c.Sex.ToString() == "0" && c.MinAge == 0 && c.MaxAge == 0 {
		return true
	}

	return false
}
