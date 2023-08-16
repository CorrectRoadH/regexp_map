package regexp_map

import (
	"regexp"
)

type RegexHashMap[T any] struct {
	internalMap map[string]T
	regexMap    map[string]T
}

func NewRegexHashMap[T any]() *RegexHashMap[T] {
	return &RegexHashMap[T]{
		internalMap: make(map[string]T),
		regexMap:    make(map[string]T),
	}
}

func (r *RegexHashMap[T]) Get(key string) (T, bool, string) {
	if value, ok := r.internalMap[key]; ok {
		return value, ok, key
	}

	for k, v := range r.regexMap {
		pattern := regexp.MustCompile(k)
		if pattern.MatchString(key) {
			return v, true, k
		}
	}

	var zero T
	return zero, false, ""
}

func (r *RegexHashMap[T]) SetStringKey(key string, value T) {
	r.internalMap[key] = value
}

func (r *RegexHashMap[T]) SetRegexpKey(key string, value T) {
	r.regexMap[key] = value
}
