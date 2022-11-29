package main

import (
	"Imp-go-topics/interfaceslearn"
	routinesdir "Imp-go-topics/routinesDir"
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
	var circle interfaceslearn.Contract1 = &interfaceslearn.Shape{Radius: 3}
	circle.Display()
	//circle.Display2()

	//-------------------example of channels-------------------

	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	go func() {
		messages <- "pong"
	}()

	msg := <-messages

	fmt.Println(msg)

	//------------------------example of channel to channel communication-------------------------------------

	pings := make(chan string, 1)

	pongs := make(chan string, 1)

	routinesdir.Ping(pings, "passed message")
	routinesdir.Pong(pings, pongs)
	fmt.Println(<-pongs)

	//-----------------------------example of wait groups--------------------------------------------
	// random goroutines are executing right now here.
	var wg1 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg1.Add(1)
		go routinesdir.Worker(i, &wg1)
	}
	//block untill the waitgroup counter goes back to 0
	//all the workers notified they're done.
	wg1.Wait()

	//-----------------------deposit and withdraw amount from account---------------------------

	var wg2 sync.WaitGroup
	wg2.Add(2)
	go routinesdir.Withdraw(700, &wg2)
	go routinesdir.Deposit(500, &wg2)
	wg2.Wait()

	fmt.Printf("New balnce %d\n", routinesdir.Balance)

	//-----------print odd and even using two different goroutines---------------

	//syncChannel := make(chan bool)

	wg3 := new(sync.WaitGroup)
	wg3.Add(2)

	go routinesdir.PrintOdd(wg3)

	go routinesdir.PrintEven(wg3)
	wg3.Wait()

}
