package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

func TestCacheMyTests(t *testing.T) {
	t.Run("MyTests", func(t *testing.T) {
		c := NewCache(3)

		require.Equal(t, 0, c.Len())

		for i := 0; i < 6; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}

		require.Equal(t, 3, c.Len())

		for i := 0; i < 3; i++ {
			val, ok := c.Get(Key(strconv.Itoa(i)))
			require.False(t, ok)
			require.Equal(t, nil, val)
		}
		for i := 3; i < 6; i++ {
			val, ok := c.Get(Key(strconv.Itoa(i)))
			require.True(t, ok)
			require.Equal(t, i, val)
		}

		c.Clear()

		require.Equal(t, 0, c.Len())

		for i := 3; i < 6; i++ {
			val, ok := c.Get(Key(strconv.Itoa(i)))
			require.False(t, ok)
			require.Equal(t, nil, val)
		}

		c.Set(Key(strconv.Itoa(6)), 6)
		val, ok := c.Get(Key(strconv.Itoa(6)))
		require.True(t, ok)
		require.Equal(t, 6, val)

		err := c.Remove(Key(strconv.Itoa(6)))
		require.True(t, err)

		val, ok = c.Get(Key(strconv.Itoa(6)))
		require.False(t, ok)

		err = c.Remove(Key(strconv.Itoa(6)))
		require.False(t, err)

		require.Equal(t, nil, val)
		require.Equal(t, 0, c.Len())
	})
}
