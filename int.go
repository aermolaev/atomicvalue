package atomicvalue

import "sync/atomic"

// A Int is an atomic int.
type Int int64

// Set64 store new int64 value.
func (i *Int) Set64(value int64) {
	atomic.StoreInt64(i.ptr(), value)
}

// Set store new int value.
func (i *Int) Set(value int) {
	i.Set64(int64(value))
}

// Get64 gets int64 value.
func (i *Int) Get64() int64 {
	return atomic.LoadInt64(i.ptr())
}

// Get gets int value.
func (i *Int) Get() int {
	return int(i.Get64())
}

// Add64 adds delta returns the new value.
func (i *Int) Add64(delta int64) int64 {
	return atomic.AddInt64(i.ptr(), delta)
}

// Add adds delta returns the new value.
func (i *Int) Add(delta int) int {
	return int(i.Add64(int64(delta)))
}

// Inc increases value by one.
func (i *Int) Inc() int {
	return i.Add(1)
}

// Dec decreases value by one.
func (i *Int) Dec() int {
	return i.Add(-1)
}

// IncDec (inc/dec)reases value by one if cond true/false.
func (i *Int) IncDec(cond bool) int {
	if cond {
		return i.Inc()
	}

	return i.Dec()
}

// Swap stores new value and returns the previous value.
func (i *Int) Swap(new int) int {
	return int(atomic.SwapInt64(i.ptr(), int64(new)))
}

// CompareAndSwap executes the compare-and-swap operation.
func (i *Int) CompareAndSwap(old, new int) bool {
	return atomic.CompareAndSwapInt64(i.ptr(), int64(old), int64(new))
}

func (i *Int) ptr() *int64 {
	return (*int64)(i)
}
