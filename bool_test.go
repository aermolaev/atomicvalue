package atomicvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolSetGet(t *testing.T) {
	assert := assert.New(t)

	var a Bool

	assert.False(a.Get())
	a.Set(true)
	assert.True(a.Get())
}

func TestBoolSetTrue(t *testing.T) {
	assert := assert.New(t)

	var a Bool

	assert.False(a.Get())
	a.SetTrue()
	assert.True(a.Get())
}

func TestBoolSetFalse(t *testing.T) {
	assert := assert.New(t)

	var a Bool

	a.SetTrue()
	assert.True(a.Get())
	a.SetFalse()
	assert.False(a.Get())
}

func TestBoolSwap(t *testing.T) {
	assert := assert.New(t)

	var a Bool

	assert.False(a.Swap(false))
	assert.False(a.Get())

	assert.False(a.Swap(true))
	assert.True(a.Get())

	assert.True(a.Swap(false))
	assert.False(a.Get())
}

func TestBoolCompareAndSwap(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		init   bool
		old    bool
		new    bool
		result bool
		store  bool
	}{
		{false, true, false, false, false},
		{false, false, true, true, true},
		{false, false, false, true, false},
		{false, true, true, false, false},
		{true, true, false, true, false},
		{true, false, true, false, true},
		{true, true, true, true, true},
		{true, false, false, false, true},
	}

	for _, tt := range tests {
		var a Bool
		a.Set(tt.init)

		r := a.CompareAndSwap(tt.old, tt.new)
		assert.Equal(tt.result, r, "result")
		assert.Equal(tt.store, a.Get(), "store")
	}
}

func BenchmarkBoolB2u(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b2u(true)
		b2u(false)
	}
}

func BenchmarkBoolGet(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.Get()
	}
}

func BenchmarkBoolSet(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.Set(true)
	}
}

func BenchmarkBoolSetTrue(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.SetTrue()
	}
}

func BenchmarkBoolSetFalse(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.SetFalse()
	}
}

func BenchmarkBoolCompareAndSwap_1(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.CompareAndSwap(false, true)
	}
}

func BenchmarkBoolCompareAndSwap_2(b *testing.B) {
	var a Bool

	for n := 0; n < b.N; n++ {
		a.CompareAndSwap(true, false)
	}
}
