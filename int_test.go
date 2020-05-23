package atomicvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSetGet(t *testing.T) {
	assert := assert.New(t)

	var a Int

	assert.Equal(0, a.Get())
	a.Set(10)
	assert.Equal(10, a.Get())
}

func TestIntSwap(t *testing.T) {
	assert := assert.New(t)

	var a Int

	assert.Equal(0, a.Swap(0))
	assert.Equal(0, a.Get())

	assert.Equal(0, a.Swap(10))
	assert.Equal(10, a.Get())

	assert.Equal(10, a.Swap(0))
	assert.Equal(0, a.Get())
}

func TestIntCompareAndSwap(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		init   int
		old    int
		new    int
		result bool
		store  int
	}{
		{0, 9, 0, false, 0},
		{0, 0, 9, true, 9},
		{0, 0, 0, true, 0},
		{0, 9, 9, false, 0},
		{9, 9, 0, true, 0},
		{9, 0, 9, false, 9},
		{9, 9, 9, true, 9},
		{9, 0, 0, false, 9},
	}

	for _, tt := range tests {
		var a Int
		a.Set(tt.init)

		r := a.CompareAndSwap(tt.old, tt.new)
		assert.Equal(tt.result, r, "result")
		assert.Equal(tt.store, a.Get(), "store")
	}
}

func BenchmarkIntGet(b *testing.B) {
	var a Int

	for n := 0; n < b.N; n++ {
		a.Get()
	}
}

func BenchmarkIntSet(b *testing.B) {
	var a Int

	for n := 0; n < b.N; n++ {
		a.Set(10)
	}
}

func BenchmarkIntCompareAndSwap_1(b *testing.B) {
	var a Int

	for n := 0; n < b.N; n++ {
		a.CompareAndSwap(0, 9)
	}
}

func BenchmarkIntCompareAndSwap_2(b *testing.B) {
	var a Int

	for n := 0; n < b.N; n++ {
		a.CompareAndSwap(9, 0)
	}
}
