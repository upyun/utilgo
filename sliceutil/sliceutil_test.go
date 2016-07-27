package sliceutil_test

import (
	"testing"
	"github.com/upyun/utilgo/sliceutil"
	"github.com/stretchr/testify/assert"
)

func TestGetSliceIndex(t testing.T) {
	s := []string{"a", "b"}

	assert := assert.New(t)

	assert.Equal(0, sliceutil.GetSliceIndex("a", s))
	assert.Equal(1, sliceutil.GetSliceIndex("b", s))
	assert.Equal(-1, sliceutil.GetSliceIndex("c", s))
}