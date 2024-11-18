package service

import (
	"easy_mall/consts"
	"time"
)

func getProductCategoryTimeRange(timestamp string) (start, end string, err error) {
	if timestamp == "" {
		start = time.Now().Format(consts.TimeFormatYearAndMoth)
		end = time.Now().Format(consts.TimeFormatYearAndMoth)
	} else {
		var t time.Time
		if t, err = time.Parse(timestamp, time.DateOnly); err != nil {
			return
		}
		start = t.Format(consts.TimeFormatYearAndMoth)
		end = t.Format(consts.TimeFormatYearAndMoth)
	}

	return
}
