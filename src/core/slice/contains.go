package slice

func ContainsUint8(s []uint8, e uint8) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
