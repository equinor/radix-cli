package time

import "time"

// Unix provides a type that can marshal and unmarshal a string representation
// of the unix epoch into a time.Time object.
type Unix struct {
	T time.Time
}
