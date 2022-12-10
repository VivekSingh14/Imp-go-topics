package main

import (
	routinesdir "Imp-go-topics/routinesDir"
	"sync"
)

var GFG = 0

func main() {
	//Modules or packages example
	//slicesandarrays.DifferencebetweenArrayandSlice()

	//--------------goroutines example-------------------
	// var w sync.WaitGroup

	// var m sync.Mutex

	// for i := 0; i < 10; i++ {
	// 	w.Add(1)
	// 	go func(wg *sync.WaitGroup, mg *sync.Mutex) {
	// 		mg.Lock()
	// 		GFG = GFG + 1
	// 		mg.Unlock()
	// 		wg.Done()
	// 	}(&w, &m)
	// }
	// w.Wait()
	// fmt.Println("Value of x", GFG)

	//----------------------interfaces and structures example-----------------------------
	// var circle interfaceslearn.Contract1 = &interfaceslearn.Shape{Radius: 3}
	// circle.Display()
	//circle.Display2()

	//-------------------example of channels-------------------

	// messages := make(chan string)

	// go func() {
	// 	messages <- "ping"
	// }()

	// go func() {
	// 	messages <- "pong"
	// }()

	// msg := <-messages

	// fmt.Println(msg)

	//------------------------example of channel to channel communication-------------------------------------

	// pings := make(chan string, 1)

	// pongs := make(chan string, 1)

	// routinesdir.Ping(pings, "passed message")
	// routinesdir.Pong(pings, pongs)
	// fmt.Println(<-pongs)

	//-----------------------------example of wait groups--------------------------------------------
	// random goroutines are executing right now here.
	// var wg1 sync.WaitGroup

	// for i := 1; i <= 5; i++ {
	// 	wg1.Add(1)
	// 	go routinesdir.Worker(i, &wg1)
	// }
	//block untill the waitgroup counter goes back to 0
	//all the workers notified they're done.
	//wg1.Wait()

	//-----------------------deposit and withdraw amount from account---------------------------

	// var wg2 sync.WaitGroup
	// wg2.Add(2)
	// go routinesdir.Withdraw(700, &wg2)
	// go routinesdir.Deposit(500, &wg2)
	// wg2.Wait()

	// fmt.Printf("New balnce %d\n", routinesdir.Balance)

	//-----------print odd and even using two different goroutines---------------

	// var wg3 sync.WaitGroup
	// m := sync.Mutex{}
	// wg3.Add(2)
	// oddch := make(chan int)
	// evench := make(chan int)

	// go routinesdir.PrintOdd(oddch)

	// go routinesdir.PrintEven(evench)

	// for i := 1; i <= 10; i++ {

	// 	if i%2 == 0 {
	// 		m.Lock()
	// 		evench <- i
	// 		<-evench
	// 		m.Unlock()
	// 		fmt.Println("mutex unlocked from even")
	// 	} else {
	// 		m.Lock()
	// 		oddch <- i
	// 		<-oddch
	// 		m.Unlock()
	// 		fmt.Println("mutex unlocked from odd")
	// 	}
	// }

	// wg3.Done()
	// wg3.Done()
	// close(oddch)
	// close(evench)

	//-----------------------url fetcher using wait group-------------------------

	var wgurl sync.WaitGroup

	urls := []string{
		"http://httpbin.org/get?x=1",
		"http://httpbin.org/get?y=2",
		"http://httpbin.org/get?z=3",
	}

	for _, u := range urls {
		wgurl.Add(1)
		go routinesdir.GetData(u, &wgurl)
	}

	wgurl.Wait()

}
