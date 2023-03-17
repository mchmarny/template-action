package cli

import (
	"github.com/mchmarny/action/pkg/action"
	"github.com/pkg/errors"
	c "github.com/urfave/cli/v2"
)

var (
	runCmd = &c.Command{
		Name:    "print",
		Aliases: []string{"p"},
		Usage:   "Prints input parameters",
		Action:  execRunCmd,
		Flags: []c.Flag{
			fileFlag,
		},
	}
)

func execRunCmd(c *c.Context) error {
	opt := &action.Options{
		File:     c.String(fileFlag.Name),
		Required: c.Bool(requiredFlag.Name),
	}

	printVersion(c)

	if err := action.Execute(c.Context, opt); err != nil {
		return errors.Wrap(err, "error executing command")
	}

	return nil
}
