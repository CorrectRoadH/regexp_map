package main

import (
	"fmt"

	"github.com/CorrectRoadH/regexp_map"
)

func main() {
	testMap := regexp_map.NewRegexHashMap[string]()
	testMap.SetStringKey("https://youtube.com", "youtube")
	testMap.SetStringKey("https://bilibili.com", "bilibili")
	testMap.SetRegexpKey("^https?://(?:www\\.)?bilibili\\.com(/[\\w-]+)*/?(\\?[^#]*)?(#.*)?$`", "bilibili")

	result1, ok, key := testMap.Get("https://youtube.com")
	fmt.Println(result1, ok, key)

	result2, ok, key := testMap.Get("https://bilibili.com")
	fmt.Println(result2, ok, key)

	result3, ok, _ := testMap.Get("https://www.bilibili.com/video/BV1394y1k7D2/")
	fmt.Println(result3, ok)

	result4, ok, key := testMap.Get("https://discord.com")
	fmt.Println(result4, ok, key)

}
