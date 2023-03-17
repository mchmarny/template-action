package cli

import (
	c "github.com/urfave/cli/v2"
)

var (
	fileFlag = &c.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "path to file",
	}

	requiredFlag = &c.BoolFlag{
		Name:    "required",
		Aliases: []string{"r"},
		Usage:   "required output",
	}
)
