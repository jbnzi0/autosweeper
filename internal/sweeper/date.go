package sweeper

import "time"

func isMoreThanOneMonthAgo(t time.Time) bool {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return t.Before(oneMonthAgo)
}
