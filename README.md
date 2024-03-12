# go-concurrent
a simple go code to execute funcs concurrent and exec func when all concurrent funcs done

## install

```shell
go get -u github.com/buddiewei/go-concurrent
```

## demo

demo ignore error
```go
package main

import (
    "fmt"
	"time"
	
    con "github.com/buddiewei/go-concurrent"
)

func main() {
    allConFuncDone := false
	cf := con.ConcurrentFunc(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
	}, func() {
		time.Sleep(3 * time.Second)
		fmt.Println("world")
	})
	cf.Add(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world")
	})
	cf.Aggregate(func() {
		fmt.Println("all done")
		allConFuncDone = true
	})
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}
```

demo sensitive to error

```go
package main

import (
	"context"
	"fmt"
	"time"
	
	con "github.com/buddiewei/go-concurrent"
)

func main() {
	allConFuncDone := false
	cf := con.ConcurrentErrFunc(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
		return nil
	}, func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("world")
		return nil
	})
	cf.Add(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world")
		return nil
	})
	err := cf.Aggregate(context.Background(), func() error {
		fmt.Println("all done")
		allConFuncDone = true
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}
```