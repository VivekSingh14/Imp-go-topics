package main

import (
	"Imp-go-topics/datastructures"
	"fmt"
	"sync"
)

var GFG = 0

func main() {
	//Modules or packages example
	//slicesandarrays.DifferencebetweenArrayandSlice()

	//goroutines example
	var w sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup, mg *sync.Mutex) {
			mg.Lock()
			GFG = GFG + 1
			mg.Unlock()
			wg.Done()
		}(&w, &m)
	}
	w.Wait()
	fmt.Println("Value of x", GFG)

	//interfaces and structures example
	var circle datastructures.Contract1 = &datastructures.Shape{Radius: 3}
	circle.Display()
	//circle.Display2()

	//print odd and even using two different goroutines

	//syncChannel := make(chan bool)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	//routinesdir.DisplayOdd(syncChannel, wg)

	//routinesdir.DisplayEven(syncChannel, wg)

}
