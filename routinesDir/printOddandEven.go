package routinesdir

import (
	"fmt"
)

func PrintEven(ch chan int) {
	for num := range ch {

		fmt.Println("Even ", num)
		//fmt.Println("even goroutine is closed")
		ch <- -1

	}

}

func PrintOdd(ch chan int) {
	for num := range ch {

		fmt.Println("Odd ", num)
		//fmt.Println("odd goroutine is closed")
		ch <- -1
	}
}
