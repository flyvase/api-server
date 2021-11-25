package response

import "harvest/src/domain/value"

type geoPoint struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func geoPointFromValue(p value.GeoPoint) geoPoint {
	return geoPoint{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}
