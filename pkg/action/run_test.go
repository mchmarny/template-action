package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestExecute(t *testing.T) {
	err := Execute(context.TODO(), nil)
	assert.Error(t, err)
	err = Execute(context.TODO(), &Options{})
	assert.Error(t, err)
	err = Execute(context.TODO(), &Options{
		File: "test",
	})
	assert.NoError(t, err)
}
