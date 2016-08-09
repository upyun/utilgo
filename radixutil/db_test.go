package radixutil_test

import (
	"fmt"
	. "testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upyun/utilgo/radixutil"
)

func TestScan(t *T) {
	keys := map[string]struct{}{}
	conf := &radixutil.ConfigMeta{
		RedisAddr:          "127.0.0.1:6379",
		RedisCluster:       false,
		RedisSentinel:      false,
		RedisSentinels:     []string{},
		RedisSentinelGroup: "master",
	}
	radixutil.InitDB(conf)
	for i := 0; i < 100; i++ {
		keys[fmt.Sprintf("scantest:%d", i)] = struct{}{}
	}

	for key := range keys {
		require.Nil(t, radixutil.Inst.Cmd("SET", key, key).Err)
	}

	output := map[string]struct{}{}
	for r := range radixutil.Inst.Scan("scantest:*") {
		output[r] = struct{}{}
	}
	assert.Equal(t, keys, output)

}
