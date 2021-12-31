package value

import (
	"api-server/src/core/slice"
	"strconv"
)

// based on ISO5218
type Sex struct {
	Code uint8
}

func NewSex(code uint8) Sex {
	if !slice.ContainsUint8([]uint8{0, 1, 2, 9}, code) {
		return Sex{
			Code: 0,
		}
	}

	return Sex{
		Code: code,
	}
}

func (s *Sex) ToString() string {
	return strconv.Itoa(int(s.Code))
}
