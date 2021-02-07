package datetime

import "time"

func Timestamp(now func() time.Time) int64 {
	return int64(time.Nanosecond) * now().UnixNano() / int64(time.Millisecond)
}
