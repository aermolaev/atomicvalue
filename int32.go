package atomicvalue

import "sync/atomic"

// A Int32 is an atomic int32.
type Int32 int32

// Set32 store new int32 value.
func (i *Int32) Set32(value int32) {
	atomic.StoreInt32(i.ptr(), value)
}

// Set store new int value.
func (i *Int32) Set(value int) {
	i.Set32(int32(value))
}

// Get32 gets int32 value.
func (i *Int32) Get32() int32 {
	return atomic.LoadInt32(i.ptr())
}

// Get64 gets int64 value.
func (i *Int32) Get64() int64 {
	return int64(i.Get32())
}

// Get gets int value.
func (i *Int32) Get() int {
	return int(i.Get32())
}

// Add32 adds delta returns the new value.
func (i *Int32) Add32(delta int32) int32 {
	return atomic.AddInt32(i.ptr(), delta)
}

// Add adds delta returns the new value.
func (i *Int32) Add(delta int) int {
	return int(i.Add32(int32(delta)))
}

// Inc increases value by one.
func (i *Int32) Inc() int {
	return i.Add(1)
}

// Dec decreases value by one.
func (i *Int32) Dec() int {
	return i.Add(-1)
}

// IncDec (inc/dec)reases value by one if cond true/false.
func (i *Int32) IncDec(cond bool) int {
	if cond {
		return i.Inc()
	}

	return i.Dec()
}

// Swap stores new value and returns the previous value.
func (i *Int32) Swap(new int) int {
	return int(atomic.SwapInt32(i.ptr(), int32(new)))
}

// CompareAndSwap executes the compare-and-swap operation.
func (i *Int32) CompareAndSwap(old, new int) bool {
	return atomic.CompareAndSwapInt32(i.ptr(), int32(old), int32(new))
}

func (i *Int32) ptr() *int32 {
	return (*int32)(i)
}
