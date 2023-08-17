# What is the Regexp Map

The Regexp Map is a specify Map. Its key can is a string or a a regexp. when the index is string. It will to index the string. If it didn't match. It will to match the regexp.

# Note
- Require Regexp is independent.
- No Thread Safe

# Usage 
```go  
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
```

# Benchmark

The time cost is O(log(n)).

```bash
➜  regexp_map git:(main) ✗ go test -bench='Bench' -benchtime 5s ./...
?       github.com/CorrectRoadH/regexp_map      [no test files]
goos: darwin
goarch: arm64
pkg: github.com/CorrectRoadH/regexp_map/benchmark
Benchmark_store_and_load_pure_Regepx_V1-10         16027            635505 ns/op
Benchmark_store_and_load_string_V1-10           18337298               395.2 ns/op
Benchmark_store_and_load_mix_V1-10                 15450            622352 ns/op
PASS
ok      github.com/CorrectRoadH/regexp_map/benchmark    35.600s
PASS
ok      github.com/CorrectRoadH/regexp_map/test 1.807s
```
# RoadMap
- [ ] add thread safe
- [ ] optimize the performance
- [ ] add Regexp check
- [ ] implment delete api
- [ ] add more usage