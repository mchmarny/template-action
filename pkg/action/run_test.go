package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	_, err := Execute(nil)
	assert.Error(t, err)
	_, err = Execute(&Request{})
	assert.Error(t, err)
	res, err := Execute(&Request{
		File: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
