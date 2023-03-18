package action

import (
	"fmt"
	"time"
)

type Response struct {
	Value       string
	ProcessedOn time.Time
}

// ProcessedOnUTC returns the processed on time in UTC.
func (r *Response) ProcessedOnUTC() string {
	if r == nil {
		return ""
	}
	return r.ProcessedOn.UTC().Format(time.RFC3339)
}

func (r *Response) String() string {
	return fmt.Sprintf("Value: %s", r.Value)
}
