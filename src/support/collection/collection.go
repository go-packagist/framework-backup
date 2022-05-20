package collection

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

type collection struct {
	items []interface{}
}

func NewCollection(items []interface{}) *collection {
	return &collection{
		items: items,
	}
}

func (c *collection) Items() []interface{} {
	return c.items
}

func (c *collection) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *collection) Remove(item interface{}) {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}

func (c *collection) Contains(item interface{}) bool {
	for _, v := range c.items {
		if v == item {
			return true
		}
	}
	return false
}

func (c *collection) Clear() {
	c.items = []interface{}{}
}

func (c *collection) Size() int {
	return len(c.items)
}

func (c *collection) IsEmpty() bool {
	return len(c.items) == 0
}

func (c *collection) Each(fn func(interface{})) {
	for _, v := range c.items {
		fn(v)
	}
}

func (c *collection) Map(fn func(interface{}) interface{}) []interface{} {
	var result []interface{}

	for _, v := range c.items {
		result = append(result, fn(v))
	}

	return result
}

func (c *collection) Filter(fn func(interface{}) bool) []interface{} {
	var result []interface{}

	for _, v := range c.items {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (c *collection) Reject(fn func(interface{}) bool) []interface{} {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

func (c *collection) Find(fn func(interface{}) bool) interface{} {
	for _, v := range c.items {
		if fn(v) {
			return v
		}
	}
	return nil
}

func (c *collection) FindIndex(fn func(interface{}) bool) int {
	for i, v := range c.items {
		if fn(v) {
			return i
		}
	}
	return -1
}

func (c *collection) FindLast(fn func(interface{}) bool) interface{} {
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i]) {
			return c.items[i]
		}
	}
	return nil
}

func (c *collection) FindLastIndex(fn func(interface{}) bool) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i]) {
			return i
		}
	}
	return -1
}

func (c *collection) Reduce(fn func(interface{}, interface{}, int) interface{}, initial interface{}) interface{} {
	var result interface{}

	result = initial

	for i, v := range c.items {
		result = fn(result, v, i)
	}

	return result
}

func (c *collection) ReduceRight(fn func(interface{}, interface{}, int) interface{}, initial interface{}) interface{} {
	var result interface{}

	result = initial

	for i := len(c.items) - 1; i >= 0; i-- {
		result = fn(result, c.items[i], i)
	}

	return result
}

func (c *collection) Every(fn func(interface{}) bool) bool {
	for _, v := range c.items {
		if !fn(v) {
			return false
		}
	}

	return true
}

func (c *collection) Some(fn func(interface{}) bool) bool {
	for _, v := range c.items {
		if fn(v) {
			return true
		}
	}

	return false
}

func (c *collection) IndexOf(item interface{}) int {
	for i, v := range c.items {
		if v == item {
			return i
		}
	}
	return -1
}

func (c *collection) LastIndexOf(item interface{}) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if c.items[i] == item {
			return i
		}
	}
	return -1
}

func (c *collection) Slice(start int, end int) []interface{} {
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

func (c *collection) SliceFrom(start int) []interface{} {
	if start < 0 {
		start = 0
	}

	return c.items[start:]
}

func (c *collection) SliceTo(end int) []interface{} {
	if end < 0 {
		end = len(c.items)
	}

	return c.items[:end]
}

func (c *collection) Reverse() {
	for i, j := 0, len(c.items)-1; i < j; i, j = i+1, j-1 {
		c.items[i], c.items[j] = c.items[j], c.items[i]
	}
}

func (c *collection) Shuffle() {
	for i := range c.items {
		j := rand.Intn(i + 1)
		c.items[i], c.items[j] = c.items[j], c.items[i]
	}
}

func (c *collection) Sort(fn func(interface{}, interface{}) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *collection) SortBy(fn func(interface{}) interface{}) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]).(int) < fn(c.items[j]).(int)
	})
}

func (c *collection) SortByDesc(fn func(interface{}) interface{}) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]).(int) > fn(c.items[j]).(int)
	})
}

func (c *collection) SortByDescFunc(fn func(interface{}, interface{}) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *collection) SortByFunc(fn func(interface{}, interface{}) bool) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	})
}

func (c *collection) Unique() []interface{} {
	var unique []interface{}

	for _, v := range c.items {
		if !c.Contains(v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (c *collection) Json() string {
	result, _ := json.Marshal(c.items)

	return string(result)
}

func (c *collection) String() string {
	return fmt.Sprintf("%v", c.items)
}

func (c *collection) First() interface{} {
	if len(c.items) > 0 {
		return c.items[0]
	}

	return nil
}

func (c *collection) Last() interface{} {
	if len(c.items) > 0 {
		return c.items[len(c.items)-1]
	}

	return nil
}
