package main

import (
	"fmt"
	"sync"
	"time"
)

func func0(index int) bool {
	fmt.Printf("\t* executing func0(%d)\n", index)
	return index < 2
}

func func1(index int) bool {
	fmt.Printf("\t* executing func1(%d)\n", index)
	return index < 2
}

func func2(index int) bool {
	fmt.Printf("\t* executing func2(%d)\n", index)
	return index < 2
}

func func3(index int) bool {
	fmt.Printf("\t* executing func3(%d)\n", index)
	return index < 2
}

func func4(index int) bool {
	fmt.Printf("\t* executing func4(%d)\n", index)
	return index < 2
}

func async(data []any) (any, error) {
	var (
		wg   sync.WaitGroup
		once sync.Once
	)
	errChan := make(chan error, 1)
	defer func() {
		if _, ok := <-errChan; ok {
			close(errChan)
		}
	}()
	for index := range data {
		println("running index", index)
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			if !func0(index) {
				once.Do(func() {
					fmt.Printf("\t# recording func0(%d)\n", index)
					errChan <- fmt.Errorf("error in func0(%d)", index)
				})
				return
			}
			if !func1(index) {
				once.Do(func() {
					fmt.Printf("\t# recording func1(%d)\n", index)
					errChan <- fmt.Errorf("error in func1(%d)", index)
				})
				return
			}
			if !func2(index) {
				once.Do(func() {
					fmt.Printf("\t# recording func2(%d)\n", index)
					errChan <- fmt.Errorf("error in func2(%d)", index)
				})
				return
			}
			if !func3(index) {
				once.Do(func() {
					fmt.Printf("\t# recording func3(%d)\n", index)
					errChan <- fmt.Errorf("error in func3(%d)", index)
				})
				return
			}
			if !func4(index) {
				once.Do(func() {
					fmt.Printf("\t# recording func4(%d)\n", index)
					errChan <- fmt.Errorf("error in func4(%d)", index)
				})
				return
			}
		}(index)
		time.Sleep(time.Microsecond * 500)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()
	if err, ok := <-errChan; ok {
		return nil, err
	}
	return len(data), nil
}

// go playground: https://go.dev/play/p/WPWzk10dnBE
func main() {
	result, err := async([]any{1, 2, 3, 4, 5})
	fmt.Printf("result: %v\nerror: %v\n", result, err)
}
