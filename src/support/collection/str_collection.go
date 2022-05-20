package collection

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

type strCollection struct {
	items []string
}

func NewStrCollection(items []string) *strCollection {
	return &strCollection{
		items: items,
	}
}

func (c *strCollection) Items() []string {
	return c.items
}

func (c *strCollection) Add(item string) *strCollection {
	c.items = append(c.items, item)

	return c
}

func (c *strCollection) Remove(item string) {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}

func (c *strCollection) RemoveAll(item string) {
	for i := 0; i < len(c.items); {
		if c.items[i] == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		} else {
			i++
		}
	}
}

func (c *strCollection) Contains(item string) bool {
	for _, v := range c.items {
		if v == item {
			return true
		}
	}

	return false
}

func (c *strCollection) Clear() {
	c.items = []string{}
}

func (c *strCollection) Size() int {
	return c.Len()
}

func (c *strCollection) Len() int {
	return len(c.items)
}

func (c *strCollection) Count() int {
	return c.Len()
}

func (c *strCollection) IsEmpty() bool {
	return len(c.items) == 0
}

func (c *strCollection) Each(fn func(string)) {
	for _, v := range c.items {
		fn(v)
	}
}

func (c *strCollection) Map(fn func(string) string) []string {
	var result []string

	for _, v := range c.items {
		result = append(result, fn(v))
	}

	return result
}

func (c *strCollection) Filter(fn func(string) bool) []string {
	var result []string

	for _, v := range c.items {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (c *strCollection) Reject(fn func(string) bool) []string {
	return c.Filter(func(v string) bool {
		return !fn(v)
	})
}

func (c *strCollection) Find(fn func(string) bool) string {
	for _, v := range c.items {
		if fn(v) {
			return v
		}
	}
	return ""
}

func (c *strCollection) FindIndex(fn func(string) bool) int {
	for i, v := range c.items {
		if fn(v) {
			return i
		}
	}
	return -1
}

func (c *strCollection) FindLast(fn func(string) bool) string {
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i]) {
			return c.items[i]
		}
	}
	return ""
}

func (c *strCollection) FindLastIndex(fn func(string) bool) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i]) {
			return i
		}
	}
	return -1
}

func (c *strCollection) Reduce(fn func(string, string, int) string, initial string) string {
	var result string

	result = initial

	for i, v := range c.items {
		result = fn(result, v, i)
	}

	return result
}

func (c *strCollection) ReduceRight(fn func(string, string, int) string, initial string) string {
	var result string

	result = initial

	for i := len(c.items) - 1; i >= 0; i-- {
		result = fn(result, c.items[i], i)
	}

	return result
}

func (c *strCollection) Every(fn func(string) bool) bool {
	for _, v := range c.items {
		if !fn(v) {
			return false
		}
	}

	return true
}

func (c *strCollection) Some(fn func(string) bool) bool {
	for _, v := range c.items {
		if fn(v) {
			return true
		}
	}

	return false
}

func (c *strCollection) IndexOf(item string) int {
	for i, v := range c.items {
		if v == item {
			return i
		}
	}
	return -1
}

func (c *strCollection) LastIndexOf(item string) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if c.items[i] == item {
			return i
		}
	}
	return -1
}

func (c *strCollection) Slice(start int, end int) []string {
	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = len(c.items)
	}

	if start > end {
		start, end = end, start
	}

	return c.items[start:end]
}

func (c *strCollection) SliceFrom(start int) []string {
	if start < 0 {
		start = 0
	}

	return c.items[start:]
}

func (c *strCollection) SliceTo(end int) []string {
	if end < 0 {
		end = len(c.items)
	}

	return c.items[:end]
}

func (c *strCollection) Reverse() {
	for i, j := 0, len(c.items)-1; i < j; i, j = i+1, j-1 {
		c.items[i], c.items[j] = c.items[j], c.items[i]
	}
}

func (c *strCollection) Shuffle() {
	for i := range c.items {
		j := rand.Intn(i + 1)
		c.items[i], c.items[j] = c.items[j], c.items[i]
	}
}

func (c *strCollection) Sort(fn func(string, string) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *strCollection) SortBy(fn func(string) string) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]) < fn(c.items[j])
	})
}

func (c *strCollection) SortByDesc(fn func(string) string) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]) > fn(c.items[j])
	})
}

func (c *strCollection) SortByDescFunc(fn func(string, string) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *strCollection) SortByFunc(fn func(string, string) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *strCollection) Unique() []string {
	var unique []string

	for _, v := range c.items {
		if !c.Contains(v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (c *strCollection) Json() string {
	result, _ := json.Marshal(c.items)

	return string(result)
}

func (c *strCollection) String() string {
	return fmt.Sprintf("%v", c.items)
}

func (c *strCollection) First() string {
	if len(c.items) > 0 {
		return c.items[0]
	}

	return ""
}

func (c *strCollection) Last() string {
	if len(c.items) > 0 {
		return c.items[len(c.items)-1]
	}

	return ""
}
