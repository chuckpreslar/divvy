package divvy_test

import "testing"

import (
	"github.com/chuckpreslar/divvy"
	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	var sample = []int{1, 2, 3}
	var result, err = divvy.New(sample).Result()
	assert.Nil(t, err, "returned unexpected error from `Result`")
	assert.Equal(t, sample, result.([]int), "returned unexpected result from `Result`")
}
