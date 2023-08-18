package geo

import (
	"math/rand"
)

// Generates random latitude and longitude coordinates within valid
// Earth ranges (-90 to 90 degrees latitude, -180 to 180 degrees longitude).
func RandomCoordinates() (float64, float64) {
	latitude := randomInRange(-90, 90)
	longitude := randomInRange(-180, 180)
	return latitude, longitude
}

// Generates a random latitude coordinate within the range of -90 to 90 degrees.
func RandomLatitude() float64 {
	return randomInRange(-90, 90)
}

// Generates a random longitude coordinate within the range of -180 to 180 degrees.
func RandomLongitude() float64 {
	return randomInRange(-180, 180)
}

// Generates a random float64 number within a specified range
// (inclusive of min and max).
func randomInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
