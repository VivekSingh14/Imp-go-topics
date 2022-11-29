package routinesdir

import (
	"fmt"
	"sync"
)

var (
	counter int
	end     int
	mutex1  sync.Mutex
)

func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex1.Lock()
	if counter < end && counter%2 == 0 {
		fmt.Print("Even ", counter)
	}
	counter += counter
	mutex1.Unlock()
}

func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex1.Lock()
	if counter < end && counter%2 != 0 {
		fmt.Print("Odd ", counter)
	}
	counter += counter
	mutex1.Unlock()

}
