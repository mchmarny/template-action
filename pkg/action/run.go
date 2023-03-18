package action

import (
	"time"

	"github.com/pkg/errors"
)

var (
	ErrMissingOptionValue = errors.New("value required")
)

// Execute runs the action. Do all the work here.
func Execute(req *Request) (*Response, error) {
	if req == nil {
		return nil, errors.New("request required")
	}
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "error validating options")
	}

	res := &Response{
		Value:       req.File,
		ProcessedOn: time.Now().UTC(),
	}

	return res, nil
}
