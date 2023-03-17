package cli

import (
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	c "github.com/urfave/cli/v2"
)

const (
	name           = "action"
	metaKeyVersion = "version"
	metaKeyCommit  = "commit"
	metaKeyDate    = "date"
)

func Execute(version string, args []string) error {
	app, err := newApp(version)
	if err != nil {
		return err
	}

	if err := app.Run(args); err != nil {
		return errors.Wrap(err, "error running app")
	}
	return nil
}

func newApp(version string) (*c.App, error) {
	if version == "" {
		return nil, errors.New("version must be set")
	}

	app := &c.App{
		EnableBashCompletion: true,
		Suggest:              true,
		Name:                 name,
		Version:              version,
		Usage:                `sample action command`,
		Compiled:             time.Now().UTC(),
		Flags: []c.Flag{
			&c.BoolFlag{
				Name:  "debug",
				Usage: "verbose output",
				Action: func(c *c.Context, debug bool) error {
					if debug {
						zerolog.SetGlobalLevel(zerolog.DebugLevel)
					}
					return nil
				},
			},
		},
		Metadata: map[string]interface{}{
			metaKeyVersion: version,
		},
		Commands: []*c.Command{
			runCmd,
		},
	}

	return app, nil
}

func printVersion(c *c.Context) {
	log.Info().Msgf(c.App.Version)
}
