package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6}

	inputCh := make(chan int)
	outputCh := make(chan int)

	go func() {
		for _, el := range numbers {
			inputCh <- el
		}
		close(inputCh)
	}()

	go func() {
		for num := range inputCh {
			outputCh <- num * 2
		}
		close(outputCh)
	}()

	for n := range outputCh {
		fmt.Println(n)
	}
}
