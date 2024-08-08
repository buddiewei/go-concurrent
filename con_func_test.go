package go_concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestConFunc(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentFunc(func() {
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

func TestConFuncWithLimit(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentFunc(func() {
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
	cf.Add(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("out range")
		s := []string{"a", "b"}
		fmt.Println(s[3])
	})
	cf.Add(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("foo")
	})
	cf.AggregateWithLimit(func() {
		fmt.Println("all done")
		allConFuncDone = true
	}, 3)
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}
