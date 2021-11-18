package entity

type Space struct {
	SpaceId             uint32
	Headline            string
	Access              string
	WeeklyVisitors      uint32
	MainCustomersSex    string
	MinMainCustomersAge uint8
	MaxMainCustomersAge uint8
	DailyPrice          uint32
	WebsiteUrl          string
	Latitude            float32
	Longitude           float32
}
