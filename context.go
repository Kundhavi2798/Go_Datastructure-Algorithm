package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//deadline := time.Now().Add(2 * time.Second)
	cxt := context.WithValue(context.Background(), "key", "value")
	//defer cancel()

	//go func() {
	//	select {
	//	case <-time.After(3 * time.Second):
	//		fmt.Println("with time cancel")
	//	case <-cxt.Done():
	//		fmt.Println("context done")
	//	}
	//}()
	go func(ctx context.Context) {
		if v := ctx.Value("key"); v != nil {
			fmt.Println("Value found:", v)
		} else {
			fmt.Println("No value found")
		}
	}(cxt)
	time.Sleep(4 * time.Second)
	//cancel()
	//time.Sleep(2 * time.Second)
}
