package core

type Sex string

// Based on ISO5218
const (
	NotKnown      = Sex("0")
	Male          = Sex("1")
	Female        = Sex("2")
	NotApplicable = Sex("9")
)

func SexFromString(s string) Sex {
	switch s {
	case "1":
		return Male
	case "2":
		return Female
	case "9":
		return NotApplicable
	// If invalid string is provided, it will return "not known"
	default:
		return NotKnown
	}
}
