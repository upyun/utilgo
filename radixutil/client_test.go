package radixutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upyun/utilgo/radixutil"
	"github.com/upyun/utilgo/testutil"
)

func TestClient(t *testing.T) {
	c := radixutil.New("localhost:6379")
	k := testutil.RandStr()
	v := testutil.RandStr()

	assert := assert.New(t)

	r, err := c.Cmd("SET", k, v).Str()
	assert.Nil(err)
	assert.Equal("OK", r)

	r, err = c.Cmd("GET", k).Str()
	assert.Nil(err)
	assert.Equal(v, r)

	ri, err := c.Cmd("DEL", k).Int()
	assert.Nil(err)
	assert.Equal(1, ri)
}
