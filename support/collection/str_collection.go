package collection

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// strCollection is a collection of strings.
type strCollection struct {
	items []string
}

// NewStrCollection creates a new string collection.
func NewStrCollection(items []string) *strCollection {
	return &strCollection{
		items: items,
	}
}

// Items returns the items in the collection.
func (c *strCollection) Items() []string {
	return c.items
}

// Add adds an item to the collection.
func (c *strCollection) Add(item string) *strCollection {
	c.items = append(c.items, item)

	return c
}

// Remove removes the first occurrence of the given item from the collection.
func (c *strCollection) Remove(item string) {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}

// RemoveAll removes all specified items from the collection.
func (c *strCollection) RemoveAll(item string) {
	for i := 0; i < len(c.items); {
		if c.items[i] == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		} else {
			i++
		}
	}
}

// Contains returns true if the collection contains the item.
func (c *strCollection) Contains(item string) bool {
	for _, v := range c.items {
		if v == item {
			return true
		}
	}

	return false
}

// Clear removes all items from the collection.
func (c *strCollection) Clear() {
	c.items = []string{}
}

// Size returns the length of the collection, alias Len().
func (c *strCollection) Size() int {
	return c.Len()
}

// Len returns the length of the collection.
func (c *strCollection) Len() int {
	return len(c.items)
}

// Count returns the length of the collection, alias Len().
func (c *strCollection) Count() int {
	return c.Len()
}

// IsEmpty returns true if the collection is empty.
func (c *strCollection) IsEmpty() bool {
	return len(c.items) == 0
}

// Each iterates over the collection calling the callback for each item.
func (c *strCollection) Each(fn func(int, string)) {
	for i, v := range c.items {
		fn(i, v)
	}
}

// Map returns a new data with the results of calling the given function
func (c *strCollection) Map(fn func(int, string) interface{}) []interface{} {
	var result []interface{}

	for i, v := range c.items {
		result = append(result, fn(i, v))
	}

	return result
}

// Filter returns a new collection with items that return true for the given function.
func (c *strCollection) Filter(fn func(int, string) bool) *strCollection {
	var result []string

	for i, v := range c.items {
		if fn(i, v) {
			result = append(result, v)
		}
	}

	return NewStrCollection(result)
}

// Reject returns a new collection with all items that do not match the given
func (c *strCollection) Reject(fn func(int, string) bool) *strCollection {
	return c.Filter(func(i int, v string) bool {
		return !fn(i, v)
	})
}

// Find returns the first item that satisfies the given predicate.
func (c *strCollection) Find(fn func(int, string) bool) (int, string) {
	for i, v := range c.items {
		if fn(i, v) {
			return i, v
		}
	}

	return -1, ""
}

// FindIndex returns the index of the first item that satisfies the given predicate.
func (c *strCollection) FindIndex(fn func(int, string) bool) int {
	i, _ := c.Find(fn)

	return i
}

// FindLast returns the last item that satisfies the given predicate.
func (c *strCollection) FindLast(fn func(int, string) bool) (int, string) {
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(i, c.items[i]) {
			return i, c.items[i]
		}
	}

	return -1, ""
}

// FindLastIndex returns the index of the last item that satisfies the given predicate.
func (c *strCollection) FindLastIndex(fn func(int, string) bool) int {
	i, _ := c.FindLast(fn)

	return i
}

// Reduce applies a function against an accumulator and each item in the collection (from left-to-right) to reduce it to a single value.
func (c *strCollection) Reduce(fn func(int, string, interface{}) interface{}, initial interface{}) interface{} {
	var result interface{}

	result = initial

	for i, v := range c.items {
		result = fn(i, v, result)
	}

	return result
}

// ReduceRight applies a function against an accumulator and each item in the collection (from right-to-left) to reduce it to a single value.
func (c *strCollection) ReduceRight(fn func(int, string, interface{}) interface{}, initial interface{}) interface{} {
	var result interface{}

	result = initial

	for i := len(c.items) - 1; i >= 0; i-- {
		result = fn(i, c.items[i], result)
	}

	return result
}

// Every returns true if all items in the collection return true for the given predicate.
func (c *strCollection) Every(fn func(int, string) bool) bool {
	for i, v := range c.items {
		if !fn(i, v) {
			return false
		}
	}

	return true
}

// Some returns true if any item in the collection returns true for the given predicate.
func (c *strCollection) Some(fn func(int, string) bool) bool {
	for i, v := range c.items {
		if fn(i, v) {
			return true
		}
	}

	return false
}

// IndexOf returns the index of the first item that matches the given value.
func (c *strCollection) IndexOf(item string) int {
	for i, v := range c.items {
		if v == item {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index of the last item that matches the given value.
func (c *strCollection) LastIndexOf(item string) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if c.items[i] == item {
			return i
		}
	}
	return -1
}

// Slice returns a new collection with the given range of items.
func (c *strCollection) Slice(start int, end int) *strCollection {
	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = len(c.items)
	}

	if start > end {
		start, end = end, start
	}

	return NewStrCollection(c.items[start:end])
}

// SliceFrom returns a new collection with the given range of items (including the start).
func (c *strCollection) SliceFrom(start int) *strCollection {
	if start < 0 {
		start = 0
	}

	return NewStrCollection(c.items[start:])
}

// SliceTo returns a new collection with the given range of items (excluding the end).
func (c *strCollection) SliceTo(end int) *strCollection {
	if end < 0 {
		end = len(c.items)
	}

	return NewStrCollection(c.items[:end])
}

// Copy returns a new collection with a copy of the items.
func (c *strCollection) Copy() *strCollection {
	cp := make([]string, len(c.items))

	copy(cp, c.items)

	return NewStrCollection(cp)
}

// Reverse returns a new collection with the items in reverse order.
func (c *strCollection) Reverse() *strCollection {
	cp := c.Copy()

	for i, j := 0, len(cp.items)-1; i < j; i, j = i+1, j-1 {
		cp.items[i], cp.items[j] = cp.items[j], cp.items[i]
	}

	return cp
}

// Shuffle returns a new collection with the items in random order.
func (c *strCollection) Shuffle() *strCollection {
	cp := c.Copy()

	for i := range cp.items {
		rand.Seed(time.Now().Unix())
		j := rand.Intn(i + 1)

		cp.items[i], cp.items[j] = cp.items[j], cp.items[i]
	}

	return cp
}

// Sort returns a new collection with the items sorted.
//	c := collection.NewStrCollection([]string{
//		"a", "b", "c", "a",
//	})
//
//	// example1:
//	c.Sort() // []string{"a", "a", "b", "c"}
//
//	// example2:
//	c2.Sort(func(a, b string) bool {
//		return a > b
//	}) // []string{"c", "b", "a", "a"}
func (c *strCollection) Sort(fn ...func(string, string) bool) {
	if len(fn) == 0 {
		sort.Strings(c.items)
	} else {
		sort.Slice(c.items, func(i, j int) bool {
			return fn[0](c.items[i], c.items[j])
		})
	}
}

// SortBy returns a new collection with the items sorted by the given key.
func (c *strCollection) SortBy(fn func(string) string) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]) < fn(c.items[j])
	})
}

// SortByDesc returns a new collection with the items sorted by the given key in descending order.
func (c *strCollection) SortByDesc(fn func(string) string) {
	sort.Slice(c.items, func(i, j int) bool {
		return fn(c.items[i]) > fn(c.items[j])
	})
}

// Unique returns a new collection with the duplicate items removed.
func (c *strCollection) Unique() *strCollection {
	unique := NewStrCollection([]string{})

	for _, v := range c.items {
		if !unique.Contains(v) {
			unique.Add(v)
		}
	}

	return unique
}

// MustJson returns the json representation of the collection.
func (c *strCollection) MustJson() string {
	result, _ := json.Marshal(c.items)

	return string(result)
}

// String returns the json representation of the collection.
func (c *strCollection) String() string {
	return fmt.Sprintf("%v", c.items)
}

// First returns the first item in the collection.
func (c *strCollection) First() string {
	if len(c.items) > 0 {
		return c.items[0]
	}

	return ""
}

// Last returns the last item in the collection.
func (c *strCollection) Last() string {
	if len(c.items) > 0 {
		return c.items[len(c.items)-1]
	}

	return ""
}
