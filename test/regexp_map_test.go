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
