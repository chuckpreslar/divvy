package divvy_test

import (
	"testing"
)

import (
	"github.com/chuckpreslar/divvy"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	var (
		sample   = []int{1, 2, 3}
		expected = []int{1, 2, 3, 4}
	)
	var result, err = divvy.New(sample).Append(4).Result()
	assert.Nil(t, err, "returned unexpected error from `Result` with `Append`")
	assert.Equal(t, expected, result.([]int), "returned unexpected result from `Result` with `Append`")
}

func TestPrepend(t *testing.T) {
	var (
		sample   = []int{2, 3, 4}
		expected = []int{1, 2, 3, 4}
	)
	var result, err = divvy.New(sample).Prepend(1).Result()
	assert.Nil(t, err, "returned unexpected error from `Result` with `Append`")
	assert.Equal(t, expected, result.([]int), "returned unexpected result from `Result` with `Append`")
}
