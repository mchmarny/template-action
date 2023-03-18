package action

import "fmt"

type Request struct {
	File     string
	Required bool
}

func (r *Request) Validate() error {
	if r.File == "" {
		return ErrMissingOptionValue
	}
	return nil
}

func (r *Request) String() string {
	return fmt.Sprintf("File: %s, Required: %t", r.File, r.Required)
}
