package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Find(t *testing.T) {
	needle := "5"
	haystack := []string{"1", "5"}

	found, loc := Find(needle, haystack)
	assert := assert.New(t)
	assert.True(found)
	assert.NotEqual(-1, loc)
}
