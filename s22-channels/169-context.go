package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Println("-----------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Println(cancel)
	fmt.Printf("%T\n", cancel)
	fmt.Println("-----------")

	cancel()
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Println(cancel)
	fmt.Printf("%T\n", cancel)
	fmt.Println("-----------")

	foo()
}

func foo() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1", ctx.Err())
	fmt.Println(runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2", ctx.Err())
	fmt.Println(runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("Cancelled Context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3:", ctx.Err())
	fmt.Println(runtime.NumGoroutine())
}
