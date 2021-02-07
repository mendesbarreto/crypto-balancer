package datetime

import (
	"testing"
	"time"
)

func TestTimestampReturnsCorrectValue(test *testing.T) {
	timeUtc := time.UTC
	nowFunction := func() time.Time {
		return time.Date(2020, 1, 1, 0, 0, 0, 0, timeUtc)
	}

	timestampExpected := int64(time.Nanosecond) * nowFunction().UnixNano() / int64(time.Millisecond)
	timestamp := Timestamp(nowFunction)

	if timestamp != timestampExpected {
		test.Error("The timestamp generated is not matching with expected")
	}
}
