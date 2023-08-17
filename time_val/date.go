package time_val

import (
	"math/rand"
	"time"
)

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
