package action

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var (
	ErrMissingOptionValue = errors.New("value required")
)

type Options struct {
	File     string
	Required bool
}

func (o *Options) Validate() error {
	if o.File == "" {
		return ErrMissingOptionValue
	}
	return nil
}

func (o *Options) String() string {
	return fmt.Sprintf("File: %s, Verbose: %t", o.File, o.Required)
}

func Execute(ctx context.Context, opt *Options) error {
	if opt == nil {
		return errors.New("options required")
	}
	if err := opt.Validate(); err != nil {
		return errors.Wrap(err, "error validating options")
	}

	log.Debug().Msgf("Options: %s", opt.String())

	return nil
}
