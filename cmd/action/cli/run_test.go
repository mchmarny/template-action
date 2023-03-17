package cli

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestImport(t *testing.T) {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	c := cli.NewContext(newTestApp(t), set, nil)
	err := execRunCmd(c)
	assert.Error(t, err)
	// TODO: add more tests
}

func newTestApp(t *testing.T) *cli.App {
	app, err := newApp("v0.0.0-test")
	assert.NoError(t, err)
	return app
}
