package regexp_map_test

import (
	"fmt"
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

func TestExample(t *testing.T) {

	reMap := regexp_map.Map[string]{}

	reMap.Store("https://youtube.com", "youtube")
	reMap.Store("https://bilibili.com", "bilibili")
	reMap.StoreRegex("^https?://(?:www\\.)?bilibili\\.com(/[\\w-]+)*/?(\\?[^#]*)?(#.*)?$", "bilibili")

	result1, ok, key := reMap.Load("https://youtube.com")
	fmt.Println(result1, ok, key) // youtube true https://youtube.com

	result2, ok, _ := reMap.Load("https://bilibili.com")
	fmt.Println(result2, ok) // bilibili true

	result3, ok, key := reMap.Load("https://www.bilibili.com/video/BV1394y1k7D2/")
	fmt.Println(result3, ok, key) // bilibili true ^https?://(?:www\.)?bilibili\.com(/[\w-]+)*/?(\?[^#]*)?(#.*)?$

	result4, ok, key := reMap.Load("https://discord.com")
	assert.False(t, ok)
	assert.Equal(t, "", result4)
	assert.Equal(t, "", key)
}
