// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {

	services := []func(ctx context.Context) (string, error){
		func(ctx context.Context) (string, error) {
			fmt.Println("Service 1 started")
			time.Sleep(2 * time.Second)
			return "", errors.New("service 1 failed")
		},
		func(ctx context.Context) (string, error) {
			fmt.Println("Service 2 started")
			time.Sleep(1 * time.Second)
			return "Service 2 Success", nil
		},
		func(ctx context.Context) (string, error) {
			fmt.Println("Service 3 started")
			time.Sleep(3 * time.Second)
			return "Service 3 Success", nil
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resch := make(chan string)
	errch := make(chan error)
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go func(svc func(ctx context.Context) (string, error)) {
			defer wg.Done()
			res, err := svc(ctx)
			if err != nil {
				errch <- err
				return
			}
			select {
			case resch <- res:
				cancel()
			case <-ctx.Done():
				return
			}

		}(service)
	}
	go func() {
		wg.Wait()
		close(resch)
		close(errch)
	}()
	count := 0
	for {
		select {
		case res, ok := <-resch:
			if ok {
				fmt.Println("SUCCESS", res)
				return
			}
		case err, ok := <-errch:
			if ok {
				fmt.Println("ERROR", err)
				count++
				return
			}

		}
	}
	fmt.Print(count)

}
