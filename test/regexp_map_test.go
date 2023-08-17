package regexp_map_test

import (
	"testing"

	"github.com/CorrectRoadH/regexp_map"
	"github.com/stretchr/testify/assert"
)

func TestStringKeyStoreAndLoad(t *testing.T) {
	testMap := regexp_map.Map[string]{}
	testMap.Store("test", "test")
	testMap.Store("1234", "41234")

	result1, ok, _ := testMap.Load("test")
	assert.True(t, ok)
	assert.Equal(t, "test", result1)

	result2, ok, _ := testMap.Load("1234")
	assert.True(t, ok)
	assert.Equal(t, "41234", result2)

	result3, ok, _ := testMap.Load("123")
	assert.False(t, ok)
	assert.Equal(t, "", result3)
}

func TestRegexpKeyStoreAndLoad(t *testing.T) {
	testMap := regexp_map.Map[string]{}
	testMap.StoreRegex("\\.cpp$", "cpp")
	testMap.StoreRegex("\\.cs$", "cs")
	testMap.StoreRegex("\\.cpp$", "c++")

	result1, ok, _ := testMap.Load("test.cpp")
	assert.True(t, ok)
	assert.Equal(t, "c++", result1)

	result4, ok, _ := testMap.Load("main.cs")
	assert.True(t, ok)
	assert.Equal(t, "cs", result4)

	result5, ok, _ := testMap.Load("hell_world.exe")
	assert.False(t, ok)
	assert.Equal(t, "", result5)
}

func TestMixKeyStoreAndLoad(t *testing.T) {
	testMap := regexp_map.Map[string]{}
	testMap.Store("test", "test")
	testMap.Store("1234", "41234")
	testMap.StoreRegex("\\.cpp$", "cpp")
	testMap.StoreRegex("\\.cs$", "cs")
	testMap.StoreRegex("\\.cpp$", "c++")

	result1, ok, _ := testMap.Load("test")
	assert.True(t, ok)
	assert.Equal(t, "test", result1)

	result2, ok, _ := testMap.Load("1234")
	assert.True(t, ok)
	assert.Equal(t, "41234", result2)

	result3, ok, _ := testMap.Load("out")
	assert.False(t, ok)
	assert.Equal(t, "", result3)

	result4, ok, _ := testMap.Load("test.cpp")
	assert.True(t, ok)
	assert.Equal(t, "c++", result4)

	result5, ok, _ := testMap.Load("main.cs")
	assert.True(t, ok)
	assert.Equal(t, "cs", result5)

	result6, ok, _ := testMap.Load("hell_world.exe")
	assert.False(t, ok)
	assert.Equal(t, "", result6)
}

func TestMoreRegexpKeyStoreAndLoad(t *testing.T) {
	testMap := regexp_map.Map[string]{}
	testMap.StoreRegex("\\.cpp$", "cpp")
	testMap.StoreRegex("\\.cs$", "cs")
	testMap.StoreRegex(`\.cpp$`, "c++")
	testMap.StoreRegex(`(bilibili|youtube)`, "video")
	testMap.StoreRegex(`(github\.com|v2ex\.com)`, "code")

	result1, ok, _ := testMap.Load("test.cpp")
	assert.True(t, ok)
	assert.Equal(t, "c++", result1)

	result2, ok, _ := testMap.Load("www.bilibili.com")
	assert.True(t, ok)
	assert.Equal(t, "video", result2)

	result3, ok, _ := testMap.Load("www.v2ex.com")
	assert.True(t, ok)
	assert.Equal(t, "code", result3)

	result4, ok, _ := testMap.Load("main.cs")
	assert.True(t, ok)
	assert.Equal(t, "cs", result4)

	result5, ok, _ := testMap.Load("hell_world.exe")
	assert.False(t, ok)
	assert.Equal(t, "", result5)
}
