package atomicvalue

import "sync/atomic"

// A Bool is an atomic bool
type Bool uint32

const (
	atomicFalse = 0
	atomicTrue  = 1
)

// Set store new value
func (a *Bool) Set(value bool) {
	atomic.StoreUint32(a.ptr(), b2u(value))
}

// SetTrue store true value
func (a *Bool) SetTrue() {
	atomic.StoreUint32(a.ptr(), atomicTrue)
}

// SetFalse store false value
func (a *Bool) SetFalse() {
	atomic.StoreUint32(a.ptr(), atomicFalse)
}

// Get gets value
func (a *Bool) Get() bool {
	return atomic.LoadUint32(a.ptr()) == atomicTrue
}

// Swap stores new value and returns the previous one
func (a *Bool) Swap(new bool) (old bool) {
	return atomic.SwapUint32(a.ptr(), b2u(new)) == atomicTrue
}

// CompareAndSwap compares and swap value
func (a *Bool) CompareAndSwap(old, new bool) bool {
	return atomic.CompareAndSwapUint32(a.ptr(), b2u(old), b2u(new))
}

func (a *Bool) ptr() *uint32 {
	return (*uint32)(a)
}

// b2u converts bool to uint32
func b2u(val bool) (u uint32) {
	if val {
		u = atomicTrue
	}

	return
}
