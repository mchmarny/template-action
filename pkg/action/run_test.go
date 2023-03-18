package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestExecute(t *testing.T) {
	_, err := Execute(context.TODO(), nil)
	assert.Error(t, err)
	_, err = Execute(context.TODO(), &Request{})
	assert.Error(t, err)
	res, err := Execute(context.TODO(), &Request{
		File: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
