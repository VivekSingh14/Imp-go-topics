package routinesdir

import (
	"fmt"
	"sync"
)

var (
	Mutex1 sync.Mutex
)

func PrintEven(ch <-chan int) {

	fmt.Print("Even ", <-ch)

}

func PrintOdd(ch <-chan int) {

	fmt.Print("Odd ", <-ch)

}
