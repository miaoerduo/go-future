# go-future

## Install

```
$ go get -u github.com/miaoerduo/go-future
```

## Usage

```go
// create promise
promise := future.NewPromise[int]()

// get future
f1 := promise.GetFuture()
f2 := promise.GetFuture()
f3 := promise.GetFuture()
f4 := promise.GetFuture()

// set value after 800ms
go func() {
    time.Sleep(time.Millisecond * 800)
    promise.SetValue(100)
}()

var wg sync.WaitGroup

wg.Add(4)

// wait for
go func() {
    defer wg.Done()

    // wait for 500 ms
    start := time.Now()
    state := f1.WaitFor(time.Millisecond * 500)
    fmt.Printf("WaitFor Elapsed %v State %v\n", time.Since(start), state)

    // wait for 500 ms, will cost 300ms
    start = time.Now()
    state = f1.WaitFor(time.Millisecond * 500)
    fmt.Printf("WaitFor Elapsed %v State %v\n", time.Since(start), state)
}()

// wait until
go func() {
    defer wg.Done()

    // wait for 500 ms
    start := time.Now()
    state := f2.WaitUntil(time.Now().Add(time.Millisecond * 500))
    fmt.Printf("WaitUntil Elapsed %v State %v\n", time.Since(start), state)

    // wait for 500 ms, will cost 300ms
    start = time.Now()
    state = f2.WaitUntil(time.Now().Add(time.Millisecond * 500))
    fmt.Printf("WaitUntil Elapsed %v State %v\n", time.Since(start), state)
}()

// wait
go func() {
    defer wg.Done()

    start := time.Now()
    // wait the value, will cost 800ms
    f3.Wait()
    fmt.Printf("Wait Elapsed %v\n", time.Since(start))
}()

// Get
go func() {
    defer wg.Done()

    start := time.Now()
    // wait the value, will cost 800ms
    value := f4.Get()
    fmt.Printf("Get Elapsed %v Value %v\n", time.Since(start), value)
}()

wg.Wait()
```
