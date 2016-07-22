package radixutil_test

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upyun/utilgo/radixutil"
)

var (
	infoRaw string
)

func init() {
	dat, err := ioutil.ReadFile("info_string.txt")
	if err != nil {
		log.Fatal("unable to read file: ", err)
	}
	infoRaw = string(dat)
}

func TestBuildMapFromInfoString(t *testing.T) {
	trimmed := strings.TrimSpace(infoRaw)
	rmap := radixutil.BuildMapFromInfoString(trimmed)

	assert.Equal(t, 0, len(rmap["redis_version"]), "redis_version should be parsed")
}

func TestBuildAllInfoMap(t *testing.T) {
	alldata := radixutil.BuildAllInfoMap(infoRaw)
	assert.Equal(t, 0, len(alldata["CPU"]["used_cpu_sys"]), "cpu.used_cpu_sys should be parsed")
}
