# atomicvalue

Wrapper for "[sync/atomic](https://golang.org/pkg/sync/atomic/)" package.

## Installation

    go get github.com/aermolaev/atomicvalue

## Simple Example

```go
var b atomicvalue.Bool

fmt.Println(b.Get()) // false

b.Set(true)
fmt.Println(b.Get()) // true
```

```go
var i atomicvalue.Int

fmt.Println(i.Get()) // 0

i.Set(10)
fmt.Println(i.Get()) // 10

i.Inc()
fmt.Println(i.Get()) // 11

i.Add(-5)
fmt.Println(i.Get()) // 6
```
