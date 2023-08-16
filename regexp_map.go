package regexp_map

import (
	"regexp"

	"golang.org/x/exp/rand"
)

type RegexHashMap[T any] struct {
	internalMap map[string]T
	regexMap    map[string]T
}

type RegexHashMapInterface[T any] interface {
	Get(key string) (T, bool, string)
	SetStringKey(key string, value T)
	SetRegexpKey(key string, value T)
}

func NewRegexHashMap[T any]() RegexHashMapInterface[T] {
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

type RegexpNode struct {
	Left      *RegexpNode
	Right     *RegexpNode
	Regexp    *regexp.Regexp
	RegexpRaw string
}

func (r *RegexpNode) Insert(re string) {
	// fmt.Println("当前节点", r, "插入", re)
	// 如果我是空
	if r.RegexpRaw == "" {
		// fmt.Println("我是空节点，把", re, "放到我身上")
		r.RegexpRaw = re[:]
		// r.Regexp = regexp.MustCompile(re)
		return
	}

	// 随机插左右
	random := rand.Intn(2)

	// 如果我是叶子
	if r.Left == nil && r.Right == nil {
		//  把当前数据放到左边，新数据放到右边
		if random == 0 {
			// fmt.Println("把", r.RegexpRaw, "放到右边")
			r.Right = &RegexpNode{
				// Regexp:    regexp.MustCompile(r.RegexpRaw[:]),
				RegexpRaw: r.RegexpRaw[:],
			}
		} else {
			// fmt.Println("把", r.RegexpRaw, "放到左边")

			r.Left = &RegexpNode{
				// Regexp:    regexp.MustCompile(r.RegexpRaw[:]),
				RegexpRaw: r.RegexpRaw[:],
			}
		}
	}

	if random == 0 {
		if r.Left == nil {
			r.Left = &RegexpNode{
				RegexpRaw: re[:],
			}
		} else {
			r.Left.Insert(re)
		}

	} else {
		if r.Right == nil {
			r.Right = &RegexpNode{
				RegexpRaw: re[:],
			}
		} else {
			r.Right.Insert(re)
		}
	}

	// 重构自己
	if r.RegexpRaw == "" {
		r.RegexpRaw = re[:]
	} else {
		r.RegexpRaw = r.RegexpRaw + "|" + re[:]
		r.Regexp = nil // 重置正则
	}
}

func (r *RegexpNode) Find(content string) (string, bool) {

	// fmt.Println("我是", r.RegexpRaw)
	if r.Left == nil && r.Right == nil {
		// fmt.Println("是叶子")
		if r.Regexp == nil {
			r.Regexp = regexp.MustCompile(r.RegexpRaw)
		}
		if r.Regexp.MatchString(content) {
			return r.RegexpRaw, true
		}
	}

	if r.Regexp == nil {
		r.Regexp = regexp.MustCompile(r.RegexpRaw)
	}
	if !r.Regexp.MatchString(content) {
		return "", false
	}

	if r.Left != nil {
		result, ok := r.Left.Find(content)
		if ok {
			return result, ok
		}
	}

	if r.Right != nil {
		result, ok := r.Right.Find(content)
		if ok {
			return result, ok
		}
	}
	return "", false
}

type RegexHashMapV2[T any] struct {
	internalMap map[string]T
	regexMap    []string
	RegexpTree  *RegexpNode
}

func NewRegexHashMapV2[T any]() RegexHashMapInterface[T] {
	return &RegexHashMapV2[T]{
		internalMap: make(map[string]T),
		regexMap:    []string{},
		RegexpTree:  &RegexpNode{},
	}
}

func (r RegexHashMapV2[T]) SetStringKey(key string, value T) {
	r.internalMap[key] = value
}

func (r *RegexHashMapV2[T]) SetRegexpKey(key string, value T) {
	_, ok := r.internalMap[key]
	if ok {
		r.internalMap[key] = value
	} else {
		// random a string
		r.regexMap = append(r.regexMap, key)
		r.internalMap[key] = value
		r.RegexpTree.Insert(key)
	}
}

// 建个二分树????🤪
func (r *RegexHashMapV2[T]) Get(key string) (T, bool, string) {

	if value, ok := r.internalMap[key]; ok {
		return value, ok, key
	}

	// fmt.Println("开始找", r.RegexpTree)
	result, ok := r.RegexpTree.Find(key)
	if ok {
		// fmt.Println("在tree中找到了[", result, "]")
		reMapResult, ok := r.internalMap[result]
		if ok {
			// fmt.Println("map中找到了", result)
			return reMapResult, ok, result
		}
	}
	// fmt.Println("啥都没有", result)

	var zero T
	return zero, false, ""
}
