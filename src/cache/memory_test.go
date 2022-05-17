package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoryStore_SetAndGet(t *testing.T) {
	memory := NewMemoryStore()

	// Put
	memory.Put("a", "aaa", time.Second*1)

	// Get
	result := memory.Get("a")
	assert.Equal(t, "aaa", result.Value())
	assert.Equal(t, "aaa", result.Val)
	assert.False(t, result.IsError())
	assert.Equal(t, nil, result.Error())

	time.Sleep(time.Second * 2)
	assert.Equal(t, nil, memory.Get("a").Value())
	assert.True(t, memory.Get("a").IsError())
}

func TestMemoryStore_Muti(t *testing.T) {
	memory, memory2 := NewMemoryStore(), NewMemoryStore()

	memory.Put("aa", "aaa", time.Second*1)
	assert.Equal(t, "aaa", memory.Get("aa").Value())
	assert.Equal(t, nil, memory2.Get("bb").Value())

	memory2.Put("bb", "bbb", time.Second*1)
	assert.Equal(t, "bbb", memory2.Get("bb").Value())
}

func TestMemoryStore_Types(t *testing.T) {
	memory := NewMemoryStore()

	memory.Put("string", "string", time.Second*1)
	assert.Equal(t, "string", memory.Get("string").Value())

	memory.Put("int", 1, time.Second*1)
	assert.Equal(t, 1, memory.Get("int").Value())

	memory.Put("float", 1.1, time.Second*1)
	assert.Equal(t, 1.1, memory.Get("float").Value())

	memory.Put("bool", true, time.Second*1)
	assert.Equal(t, true, memory.Get("bool").Value())

	memory.Put("nil", nil, time.Second*1)
	assert.Equal(t, nil, memory.Get("nil").Value())

	type test struct {
		Name string
	}

	memory.Put("struct", test{Name: "test"}, time.Second*1)
	assert.Equal(t, test{Name: "test"}, memory.Get("struct").Value())

	memory.Put("map", map[string]string{"a": "b"}, time.Second*1)
	assert.Equal(t, map[string]string{"a": "b"}, memory.Get("map").Value())

	memory.Put("slice", []string{"a", "b"}, time.Second*1)
	assert.Equal(t, []string{"a", "b"}, memory.Get("slice").Value())
}
