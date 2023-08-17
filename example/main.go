package main

import (
	"fmt"

	"github.com/CorrectRoadH/regexp_map"
)

func main() {
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
	fmt.Println(result4, ok, key) //  false
}
