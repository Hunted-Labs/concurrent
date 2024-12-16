# concurrent

[![Sourcegraph](https://sourcegraph.com/github.com/Hunted-Labs/concurrent/-/badge.svg)](https://sourcegraph.com/github.com/Hunted-Labs/concurrent?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/Hunted-Labs/concurrent)
[![Build Status](https://travis-ci.org/Hunted-Labs/concurrent.svg?branch=master)](https://travis-ci.org/Hunted-Labs/concurrent)
[![codecov](https://codecov.io/gh/Hunted-Labs/concurrent/branch/master/graph/badge.svg)](https://codecov.io/gh/Hunted-Labs/concurrent)
[![rcard](https://goreportcard.com/badge/github.com/Hunted-Labs/concurrent)](https://goreportcard.com/report/github.com/Hunted-Labs/concurrent)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/Hunted-Labs/concurrent/master/LICENSE)

* concurrent.Map: backport sync.Map for go below 1.9
* concurrent.Executor: goroutine with explicit ownership and cancellable

# concurrent.Map

because sync.Map is only available in go 1.9, we can use concurrent.Map to make code portable

```go
m := concurrent.NewMap()
m.Store("hello", "world")
elem, found := m.Load("hello")
// elem will be "world"
// found will be true
```

# concurrent.Executor

```go
executor := concurrent.NewUnboundedExecutor()
executor.Go(func(ctx context.Context) {
    everyMillisecond := time.NewTicker(time.Millisecond)
    for {
        select {
        case <-ctx.Done():
            fmt.Println("goroutine exited")
            return
        case <-everyMillisecond.C:
            // do something
        }
    }
})
time.Sleep(time.Second)
executor.StopAndWaitForever()
fmt.Println("executor stopped")
```

attach goroutine to executor instance, so that we can

* cancel it by stop the executor with Stop/StopAndWait/StopAndWaitForever
* handle panic by callback: the default behavior will no longer crash your application