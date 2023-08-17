package geo

import (
	"math/rand"
)

func RandomCoordinates() (float64, float64) {
	latitude := rand.Float64()*(180.0*2) - 180.0
	longitude := rand.Float64()*(360.0*2) - 360.0
	return latitude, longitude
}
