package bubble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	src := []int{1, 3, 5, 4, 2, 7}
	BubbleSort[int](src)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 7}, src)
}
