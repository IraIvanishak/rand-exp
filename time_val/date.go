package time_val

import (
	"math/rand"
	"time"
)

// Function generates a random date and time within the current year
func RandomDateTime() time.Time {
	currentYear := time.Now().Year()

	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(31) + 1
	hour := rand.Intn(24)
	minute := rand.Intn(60)
	second := rand.Intn(60)

	randomDateTime := time.Date(currentYear, month, day, hour, minute, second, 0, time.UTC)
	return randomDateTime
}

// Generates a random date and time within a specified range between start and end
func RandomDateInRange(start, end time.Time) time.Time {
	delta := end.Sub(start)
	randomDelta := time.Duration(rand.Int63n(int64(delta)))
	randomTime := start.Add(randomDelta)

	randomTime = randomTime.Truncate(time.Second)

	return randomTime
}
