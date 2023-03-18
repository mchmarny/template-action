package main

import (
	"strconv"
	"time"

	"github.com/mchmarny/template-action/pkg/action"
	gha "github.com/sethvargo/go-githubactions"
)

var (
	version = "v0.0.1-default"
)

func main() {
	// init action with the version and build time
	a := gha.WithFieldsMap(map[string]string{
		"version": version,
		"build":   time.Now().UTC().Format(time.RFC3339),
	})

	// log start and end
	a.Infof("starting action")
	defer a.Infof("action completed")

	// parse optional input
	isReq, _ := strconv.ParseBool(a.GetInput("required"))
	a.Debugf("required: %t", isReq)

	// create request
	req := &action.Request{
		File:     a.GetInput("file"),
		Required: isReq,
	}

	// execute action with the request
	res, err := action.Execute(req)
	if err != nil {
		a.Fatalf("error: %s", err)
	}

	// set output
	a.SetOutput("output", res.Value)
	a.SetOutput("processed_on", res.ProcessedOnUTC())
}
