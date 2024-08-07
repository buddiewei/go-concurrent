package go_concurrent

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestConErrFunc_Nil(t *testing.T) {
	cf := ConcurrentErrFunc()
	allConFuncDone := false
	cf.Add(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
		return nil
	})
	cf.Add(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("world")
		return nil
	})
	err := cf.Aggregate(context.Background(), func() error {
		fmt.Println("all done")
		allConFuncDone = true
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}

func TestConErrFunc_NoErr(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentErrFunc(func() error {
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
		t.Fatal(err)
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}

func TestConErrFunc_Err(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentErrFunc(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
		return nil
	}, func() error {
		time.Sleep(3 * time.Second)
		return fmt.Errorf("sleep 3s error")
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
	if err == nil {
		t.Fatal(fmt.Errorf("expected error, but got nil"))
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}

func TestConErrFunc_AggregateErr(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentErrFunc(func() error {
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
		allConFuncDone = true
		fmt.Println("error is occurred when aggregating")
		return fmt.Errorf("aggregate error")
	})
	if err == nil {
		t.Fatal(fmt.Errorf("expected error, but got nil"))
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}

func TestConErrFuncWithLimit_Err(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentErrFunc(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
		return nil
	})
	cf.Add(func() error {
		time.Sleep(4 * time.Second)
		fmt.Println("out range")
		s := []string{"a", "b"}
		fmt.Println(s[3])
		return nil
	})
	cf.Add(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world")
		return nil
	})
	cf.Add(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("foo")
		return nil
	})
	err := cf.AggregateWithLimit(context.Background(), func() error {
		fmt.Println("all done")
		allConFuncDone = true
		return nil
	}, 3)
	if err == nil {
		t.Fatal(fmt.Errorf("expected error, but got nil"))
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}

func TestConErrFuncWithLimit_NoErr(t *testing.T) {
	allConFuncDone := false
	cf := ConcurrentErrFunc(func() error {
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
	cf.Add(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("foo")
		return nil
	})
	cf.Add(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("bar")
		return nil
	})
	err := cf.AggregateWithLimit(context.Background(), func() error {
		fmt.Println("all done")
		allConFuncDone = true
		return nil
	}, 2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("allConFuncDone: %t", allConFuncDone)
}
