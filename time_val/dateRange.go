package time_val

import (
	"math/rand"
	"time"
)

func RandomDateInRange(start, end time.Time) time.Time {
	delta := end.Sub(start)
	randomDelta := time.Duration(rand.Int63n(int64(delta)))
	randomTime := start.Add(randomDelta)

	randomTime = randomTime.Truncate(time.Second)

	return randomTime
}
