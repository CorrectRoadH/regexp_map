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

func init_random_regexq_data(count int) ([]string, []string) {
	regexKeys := []string{}
	regexIndexs := []string{}
	for i := 0; i < (count + 2); i++ {
		format := randomString(3)
		regexKeys = append(regexKeys, "\\."+format+"$")
		regexIndexs = append(regexIndexs, randomString(5)+"."+format)
	}
	return regexKeys, regexIndexs
}

func init_random_string_data(count int) ([]string, []string) {
	regexKeys := []string{}
	for i := 0; i < (count + 2); i++ {
		regexKeys = append(regexKeys, randomString(5))
	}
	return regexKeys, regexKeys
}

func benchmark_set_regexp_key(b *testing.B, m regexp_map.RegexpMap[string], regexKeys []string, regexIndexs []string) {
	for i := 0; i < b.N; i++ {
		m.StoreRegex(regexKeys[i], regexIndexs[i])
	}
}

func benchmark_set_string_key(b *testing.B, m regexp_map.RegexpMap[string], regexKeys []string, regexIndexs []string) {
	for i := 0; i < b.N; i++ {
		m.Store(regexKeys[i], regexIndexs[i])
	}
}

func benchmark_get_regexpKey(b *testing.B, m regexp_map.RegexpMap[string], regexIndexs []string) {
	for i := 0; i < b.N; i++ {
		_, ok, _ := m.Load(regexIndexs[i])
		if !ok {
			b.Fatal("not found")
		}
	}
}

func Benchmark_store_and_load_pure_Regepx_V1(b *testing.B) {
	k, i := init_random_regexq_data(b.N)

	temp_map := &regexp_map.Map[string]{}
	benchmark_set_regexp_key(b, temp_map, k, i)
	benchmark_get_regexpKey(b, temp_map, i)
}

func Benchmark_store_and_load_string_V1(b *testing.B) {
	k, i := init_random_string_data(b.N)

	temp_map := &regexp_map.Map[string]{}
	benchmark_set_string_key(b, temp_map, k, i)
	benchmark_get_regexpKey(b, temp_map, i)
}

func Benchmark_store_and_load_mix_V1(b *testing.B) {
	k, i := init_random_string_data(b.N)

	temp_map := &regexp_map.Map[string]{}
	benchmark_set_string_key(b, temp_map, k, i)

	k2, i2 := init_random_regexq_data(b.N)
	benchmark_set_regexp_key(b, temp_map, k2, i2)

	benchmark_get_regexpKey(b, temp_map, i)
	benchmark_get_regexpKey(b, temp_map, i2)
}
