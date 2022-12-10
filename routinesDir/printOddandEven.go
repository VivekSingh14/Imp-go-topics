package routinesdir

import (
	"fmt"
)

func PrintEven(ch chan int) {
	for {
		select {
		case eve := <-ch:
			fmt.Println("Even ", eve)
			fmt.Println("even goroutine is closed")
			ch <- -1
		}
	}

}

func PrintOdd(ch chan int) {
	for {
		select {
		case od := <-ch:
			fmt.Println("Odd ", od)
			fmt.Println("odd goroutine is closed")
			ch <- -1
		}
	}

}
