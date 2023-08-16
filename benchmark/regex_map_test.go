package benchmark

import (
	"math/rand"
	"testing"

	"github.com/CorrectRoadH/regexp_map"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var stringKeys = []string{}

func initRandomData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringKeys = append(stringKeys, randomString(5))
	}
}

func initRandomRegexqData(b *testing.B) ([]string, []string) {
	regexKeys := []string{}
	regexIndexs := []string{}
	for i := 0; i < b.N; i++ {
		format := randomString(3)
		regexKeys = append(regexKeys, "\\."+format+"$")
		regexIndexs = append(regexIndexs, randomString(5)+"."+format)
	}
	return regexKeys, regexIndexs
}

func benchmarkMapStringKey(b *testing.B, m map[string]string) {
	for i := 0; i < b.N; i++ {
		_, ok := m[stringKeys[i]]
		if !ok {
			b.Fatal("not found")
		}
	}
}

// func benchmarkRegexpMapStringKey(b *testing.B, m regexp_map.RegexHashMapInterface[string]) {
// 	for i := 0; i < b.N; i++ {
// 		m.SetStringKey(stringKeys[i], "test")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		_, ok, _ := m.Get(stringKeys[i])
// 		if !ok {
// 			b.Fatal("not found")
// 		}
// 	}

// }

// func benchmarkRegexpKey(b *testing.B, m regexp_map.RegexHashMapInterface[string]) {
// 	for i := 0; i < b.N; i++ {
// 		_, ok, _ := v2map.Get(regexIndexs[i%10000])
// 		if !ok {
// 			b.Fatal("not found")
// 		}
// 	}
// }

func benchmark_set_regexpKey(b *testing.B, m regexp_map.RegexHashMapInterface[string], regexKeys []string) {
	for i := 0; i < b.N; i++ {
		m.SetRegexpKey(regexKeys[i], "test")
	}
}

func benchmark_get_regexpKey(b *testing.B, m regexp_map.RegexHashMapInterface[string], regexIndexs []string) {
	for i := 0; i < b.N; i++ {
		_, ok, _ := m.Get(regexIndexs[i])
		if !ok {
			b.Fatal("not found")
		}
	}
}

// func Benchmark_Map(b *testing.B) {
// 	initRandomData(b)
// 	benchmarkMapStringKey(b, map[string]string{})
// }

// func Benchmark_RegexpMap(b *testing.B) {
// 	initRandomData(b)
// 	benchmarkRegexpMapStringKey(b, regexp_map.NewRegexHashMap[string]())
// }

// func Benchmark_All_RegexpMap(b *testing.B) {
// 	initRandomRegexqData(b)
// 	benchmarkAllKey(b, regexp_map.NewRegexHashMap[string]())
// }

func Benchmark_PureRegepx_V2RegexpMap(b *testing.B) {
	k, i := initRandomRegexqData(b)

	// f, _ := os.OpenFile("2-cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	tmap := regexp_map.NewRegexHashMapV2[string]()
	benchmark_set_regexpKey(b, tmap, k)
	benchmark_get_regexpKey(b, tmap, i)
}

// func Benchmark_All_V2RegexpMap(b *testing.B) {
// 	initRandomRegexqData(b)
// 	benchmarkAllKey(b, regexp_map.NewRegexHashMapV2[string]())
// }

func Benchmark_PureRegepx_RegexpMap(b *testing.B) {
	k, i := initRandomRegexqData(b)

	// f, _ := os.OpenFile("1-cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	tmap := regexp_map.NewRegexHashMap[string]()
	benchmark_set_regexpKey(b, tmap, k)
	benchmark_get_regexpKey(b, tmap, i)
}
