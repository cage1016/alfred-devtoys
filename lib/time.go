package lib

import (
	"fmt"
	"math"
	"time"
)

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	// https://www.unixtimestamp.com
	// https://www.jotform.com/help/443-mastering-date-and-time-calculation/
	Month    = 2629743  // 30.44 days
	Year     = 31556926 // 365.24 days
	LongTime = 37 * Year
)

func TimeDuration(diff time.Duration) string {
	diff /= time.Second

	lbl := "from now"

	after := diff > 0

	if after {
		lbl = "ago"
		diff *= 1
	} else {
		diff *= -1
	}

	switch {
	case diff <= 0:
		return "now"

	case diff < 1*Minute:
		return fmt.Sprintf("%dsec %s", diff, lbl)

	case diff < 1*Hour:
		return fmt.Sprintf("%dmin %dsec %s",
			int64(math.Floor(float64((diff%Day)%Hour/Minute))),
			int64(math.Mod(float64(diff), 60)),
			lbl)

	case diff < 1*Day:
		return fmt.Sprintf("%dhr %dmin %dsec %s",
			int64(math.Floor(float64(diff%Day/Hour))),
			int64(math.Floor(float64((diff%Day)%Hour/Minute))),
			int64(math.Mod(float64(diff), 60)),
			lbl)

	case diff < 1*Month:
		return fmt.Sprintf("%dd %dhr %dmin %dsec %s",
			int64(math.Floor(float64((diff%Year)%Month/Day))),
			int64(math.Floor(float64(diff%Day/Hour))),
			int64(math.Floor(float64((diff%Day)%Hour/Minute))),
			int64(math.Mod(float64(diff), 60)),
			lbl)

	case diff < 1*Year:
		return fmt.Sprintf("%dmo %dd %dhr %dmin %dsec %s",
			int64(math.Floor(float64(diff%Year/Month))),
			int64(math.Floor(float64((diff%Year)%Month/Day))),
			int64(math.Floor(float64(diff%Day/Hour))),
			int64(math.Floor(float64((diff%Day)%Hour/Minute))),
			int64(math.Mod(float64(diff), 60)),
			lbl)

	case diff < LongTime:
		return fmt.Sprintf("%dyr %dmo %dd %dhr %dmin %dsec %s",
			int64(math.Floor(float64(diff/Year))),
			int64(math.Floor(float64(diff%Year/Month))),
			int64(math.Floor(float64((diff%Year)%Month/Day))),
			int64(math.Floor(float64(diff%Day/Hour))),
			int64(math.Floor(float64((diff%Day)%Hour/Minute))),
			int64(math.Mod(float64(diff), 60)),
			lbl)
	}

	return "long ago"
}
