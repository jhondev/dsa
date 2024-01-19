package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	v, _ := SearchLinear(src, 3)
	assert.Equal(t, 2, v)
}

func TestNotFound(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	_, found := SearchLinear(src, 6)
	assert.False(t, found)
}
