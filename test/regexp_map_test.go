package regexp_map_test

import (
	"testing"

	"github.com/CorrectRoadH/regexp_map"
	"github.com/stretchr/testify/assert"
)

func TestRegexHashMap(t *testing.T) {
	testMap := regexp_map.NewRegexHashMap[string]()
	testMap.SetStringKey("test", "test")
	testMap.SetStringKey("1234", "41234")
	testMap.SetRegexpKey("\\.cpp$", "41234")

	result1, ok, _ := testMap.Get("test")
	assert.True(t, ok)
	assert.Equal(t, "test", result1)

	result2, ok, _ := testMap.Get("hello.cpp")
	assert.True(t, ok)
	assert.Equal(t, "41234", result2)

	result3, ok, _ := testMap.Get("\\.cs$")
	assert.False(t, ok)
	assert.Equal(t, "", result3)
}

func TestRegexV2HashMap(t *testing.T) {
	testMap := regexp_map.NewRegexHashMapV2[string]()
	testMap.SetStringKey("test", "test")
	testMap.SetStringKey("1234", "41234")
	testMap.SetRegexpKey("\\.cpp$", "cpp")
	testMap.SetRegexpKey("\\.tpp$", "tpp")
	testMap.SetRegexpKey("\\.dpp$", "dpp")
	testMap.SetRegexpKey("\\.gpp$", "gpp")
	testMap.SetRegexpKey("\\.hpp$", "hpp")
	testMap.SetRegexpKey("\\.fpp$", "fpp")
	testMap.SetRegexpKey("\\.g34pp$", "fuck")
	testMap.SetRegexpKey("\\.hpp$", "hpp")
	testMap.SetRegexpKey("\\.fpp$", "fpp")
	testMap.SetRegexpKey("\\.g34pp$", "fucfk")
	testMap.SetRegexpKey("\\.hp234p$", "hpp")
	testMap.SetRegexpKey("\\.fp234p$", "fpp")
	testMap.SetRegexpKey("\\.g34324pp$", "fuck")
	testMap.SetRegexpKey("\\.g34pp$", "ddd")
	testMap.SetRegexpKey("\\.fhpp$", "hpp")
	testMap.SetRegexpKey("\\.ffpp$", "fpp")
	testMap.SetRegexpKey("\\.g3234pp$", "fuck")
	testMap.SetRegexpKey("\\.hp223134p$", "hpp")
	testMap.SetRegexpKey("\\.fp234pgf$", "fpp")
	testMap.SetRegexpKey("\\.g34324sdpp$", "fuck")
	testMap.SetRegexpKey("\\.g34psdfp$", "fuck")
	testMap.SetRegexpKey("\\.hpsdsdaffp$", "hpp")
	testMap.SetRegexpKey("\\.fpssdfdfp$", "fpp")
	testMap.SetRegexpKey("\\.g34psdfp$", "fuck")
	testMap.SetRegexpKey("\\.hp234dsafp$", "hpp")
	testMap.SetRegexpKey("\\.fp234sdasdffp$", "fpp")
	testMap.SetRegexpKey("\\.g34324pdsafsdp$", "fuck")

	result1, ok, _ := testMap.Get("test")
	assert.True(t, ok)
	assert.Equal(t, "test", result1)

	result2, ok, _ := testMap.Get("h313el423lo.cpp")
	assert.True(t, ok)
	assert.Equal(t, "cpp", result2)

	result4, ok, _ := testMap.Get("he2134llo.dpp")
	assert.True(t, ok)
	assert.Equal(t, "dpp", result4)

	result3, ok, _ := testMap.Get("\\.cs$")
	assert.False(t, ok)
	assert.Equal(t, "", result3)

	result5, ok, _ := testMap.Get("\\.c2134s$")
	assert.False(t, ok)
	assert.Equal(t, "", result5)

	result6, ok, _ := testMap.Get("test.cpp.cpp")
	assert.True(t, ok)
	assert.Equal(t, "cpp", result6)

	result7, ok, _ := testMap.Get("test.cpp.cpp.g34pp")
	assert.True(t, ok)
	assert.Equal(t, "ddd", result7)

	result8, ok, _ := testMap.Get("test.g34324pdsafsdp")
	assert.True(t, ok)
	assert.Equal(t, "fuck", result8)

}
