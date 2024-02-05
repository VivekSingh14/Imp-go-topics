package main

import (
	"fmt"
	"strings"
)

var GFG = 0

// func odd1(od chan int) {
// 	for {
// 		fmt.Println(<-od)
// 	}

// }

// func even1(ev chan int) {

// 	for {
// 		fmt.Println(<-ev)
// 	}
// }

//func main() {
// 	//Modules or packages example
//slicesandarrays.DifferencebetweenArrayandSlice()

// 	//--------------goroutines example-------------------
// var w sync.WaitGroup

// var m sync.Mutex

// for i := 0; i < 10; i++ {
// 	w.Add(1)
// 	go func(wg *sync.WaitGroup, mg *sync.Mutex) { //we have defined the function with argument
// 		mg.Lock()
// 		GFG = GFG + 1 //this method definition
// 		mg.Unlock()
// 		wg.Done()
// 	}(&w, &m) //here we called the function and passed the argument
// }
// w.Wait()
// fmt.Println("Value of x", GFG)

// 	//----------------------interfaces and structures example-----------------------------
// var circle interfaceslearn.Contract1 = &interfaceslearn.Shape{Radius: 3}
// circle.Display()
//circle.Display2()

// 	//-------------------example of channels-------------------

// messages := make(chan string)

// go func() {
// 	messages <- "ping"

// }()

// go func() {
// 	messages <- "pong"

// }()

// msg := <-messages

// fmt.Println(msg)

// close(messages)

// 	//------------------------example of channel to channel communication-------------------------------------

// pings := make(chan string, 1)

// pongs := make(chan string, 1)

// routinesdir.Ping(pings, "passed message")
// routinesdir.Pong(pings, pongs)
// fmt.Println(<-pongs)

// 	//-----------------------------example of wait groups--------------------------------------------
// 	// random goroutines are executing right now here.
// var wg1 sync.WaitGroup

// for i := 1; i <= 5; i++ {
// 	wg1.Add(1)
// 	go routinesdir.Worker(i, &wg1)
// }
// //block untill the waitgroup counter goes back to 0
// //all the workers notified they're done.
// wg1.Wait()

// 	//-----------------------deposit and withdraw amount from account---------------------------

// var wg2 sync.WaitGroup
// wg2.Add(2)
// go routinesdir.Withdraw(700, &wg2)
// go routinesdir.Deposit(500, &wg2)
// wg2.Wait()
// fmt.Printf("New balnce %d\n", routinesdir.Balance)

// 	//-----------print odd and even using two different goroutines---------------

// oddch := make(chan int)
// evench := make(chan int)

// go routinesdir.PrintOdd(oddch)

// go routinesdir.PrintEven(evench)

// for i := 1; i <= 10; i++ {

// 	if i%2 == 0 {
// 		evench <- i
// 		//fmt.Println("even....")
// 		<-evench
// 	} else {
// 		//m.Lock()
// 		oddch <- i
// 		<-oddch
// 		//fmt.Println("odd....")
// 	}
// 	//fmt.Println("looping...")

// }

// close(oddch)
// close(evench)

// 	//-----------------------url fetcher using wait group-------------------------

// 	// var wgurl sync.WaitGroup

// 	// urls := []string{
// 	// 	"http://httpbin.org/get?x=1",
// 	// 	"http://httpbin.org/get?y=2",
// 	// 	"http://httpbin.org/get?z=3",
// 	// }

// 	// for _, u := range urls {
// 	// 	wgurl.Add(1)
// 	// 	go routinesdir.GetData(u, &wgurl)
// 	// }

// 	// wgurl.Wait()
// 	//----------------------------trying to pass one data to 2 channels same time--------------

// 	// tings := make(chan string, 1)
// 	// tings <- "my name is vivek"
// 	// go routinesdir.Tong(tings)
// 	// go routinesdir.Ting(tings)

// 	// time.Sleep(2 * time.Second)

// 	// ---------------------------buffered channels --------------------------------------
// 	// ch := make(chan string, 2)
// 	// ch <- "My name is anthony and i am being passed through the channel 1st time"
// 	// ch <- "My name is anthony and i am being passed through the channel 2nd time"
// 	// ch <- "My name is anthony and i am being passed through the channel 3rd time"
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)

// 	// --------------------------buffered channel int ------------------------------------

// 	// ch := make(chan int, 2)
// 	// 	// go routinesdir.Write(ch)
// 	// 	// time.Sleep(2 * time.Second)

// 	// 	// for v := range ch {
// 	// 	// 	fmt.Println("read value ", v, " from ch")
// 	// 	// 	time.Sleep(2 * time.Second)
// 	// 	// }

// 	//new and optimized way to print odd and even sequentially

// 	// even := make(chan int)
// 	// odd := make(chan int)

// 	// go odd1(odd)
// 	// go even1(even)

// 	// for i := 1; i <= 10; i++ {
// 	// 	if i%2 == 0 {
// 	// 		even <- i

// 	// 	} else {
// 	// 		odd <- i
// 	// 	}
// 	// }

// 	ch := make(chan string)

// 	for i := 0; i < 10; i++ {
// 		go temp(i, ch)
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i, " ", <-ch)
// 	}

// }

// func temp(i int, ch chan string) {
// 	ch <- "Goroutines: " + strconv.Itoa(i)
// }

// func generateNumbers(total int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for idx := 1; idx <= total; idx++ {
// 		fmt.Printf("Generating number %d\n", idx)
// 	}
//}

//---------------------------------------------------------------------------//

/*

   Goroutines:

   Goroutines are lightweight threads of execution that are managed by the Go runtime.
   They are a key feature of Go that allows for concurrent programming.
   In this program, we use two goroutines: the main goroutine and the printNumbers goroutine.
   The main goroutine is responsible for sending values to the channel chan1, while the printNumbers goroutine is responsible for receiving values from the channel and printing them to the console.

   Channels:

   Channels are a communication mechanism for goroutines.
   They allow goroutines to send and receive values to each other.
   In this program, we use a channel to communicate between the main goroutine and the printNumbers goroutine.
   The main goroutine sends values to the channel, and the printNumbers goroutine receives them.

   Wait Groups:

   Wait groups are a synchronization mechanism for goroutines.
   They allow the main goroutine to wait for all of the other goroutines to finish before exiting.
   In this program, we use a wait group to wait for the printNumbers goroutine to finish before exiting the main goroutine.

   Select Statement:

   The select statement is a powerful feature of Go that allows goroutines to wait for multiple events to occur.
   In this program, we use the select statement to allow the printNumbers goroutine to wait for either a value to be sent to the channel or for the done signal to be received.

   Done Channel:

   The done channel is a special channel that is used to signal to a goroutine that it should exit.
   In this program, we use the done channel to signal to the printNumbers goroutine that the main goroutine is finished sending values.

   Closing Channels:

   Closing a channel indicates that no more values will be sent to it.
   In this program, we close the channel chan1 after the main goroutine has sent all of the values to it.
   This signals to the printNumbers goroutine that it should exit.



*/

// func printNumbers(wg *sync.WaitGroup, chan1 chan int, done <-chan bool) {

// 	for {
// 		select {
// 		case <-done:
// 			wg.Done()
// 			return
// 		case ch := <-chan1:
// 			fmt.Println(ch)

// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	chan1 := make(chan int)
// 	done := make(chan bool)

// 	wg.Add(1)
// 	go printNumbers(&wg, chan1, done)

// 	for i := 1; i <= 10; i++ {
// 		chan1 <- i
// 	}

// 	done <- true

// 	close(chan1)

// 	fmt.Println("Waiting for goroutines to finish...")
// 	wg.Wait()
// 	fmt.Println("Done!")
// }

// func myfunc(ch chan int) {

// 	fmt.Println(234 + <-ch)

// }
// func main() {
// 	fmt.Println("start Main method")
// 	// Creating a channel
// 	ch := make(chan int)

// 	go myfunc(ch)

// 	ch <- 23

// 	fmt.Println("End Main method")
// }

// func processChunk(data []int, ch chan int) {
// 	sum := 0
// 	for _, num := range data {
// 		sum += num
// 	}
// 	ch <- sum
// }

// func main() {
// 	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	ch := make(chan int)
// 	wg := sync.WaitGroup{}

// 	chunkSize := 3
// 	for i := 0; i < len(data) && i+chunkSize <= len(data); i += chunkSize {
// 		chunk := data[i : i+chunkSize]
// 		wg.Add(1)

// 		go processChunk(chunk, ch)
// 	}

// 	totalSum := 0
// 	for sum := range ch {
// 		totalSum += sum
// 	}

// 	wg.Wait()
// 	close(ch)

// 	fmt.Println("Total sum of the dataset:", totalSum)
// }

/*
   New example of goroutine and channels
*/

// func main() {
// 	ch := make(chan int)
// 	go printNum(ch)
// 	//in this example we are receiving the 2 extra zeros, because channel closed after 7th iteration and channel's default value would be zero
// 	for i := 1; i < 4; i++ {
// 		val, ok := <-ch
// 		if !ok {
// 			fmt.Println("closed channel value: ", val)
// 		} else {
// 			fmt.Println(val)
// 		}
// 		if i == 2 {
// 			close(ch)
// 		}
// 	}

// 	//close(ch)
// }

// func printNum(cha chan int) {
// 	for i := 1; i < 4; i++ {
// 		cha <- i
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	var odd_ch = make(chan int)
// 	var even_ch = make(chan int)
// 	var flag = make(chan bool)

// 	wg.Add(2)

// 	go print_even(even_ch, flag, &wg)
// 	go print_odd(odd_ch, flag, &wg)

// 	for i := 1; i <= 12; i++ {
// 		if i%2 == 0 {
// 			even_ch <- i
// 		} else {
// 			odd_ch <- i
// 		}
// 		<-flag
// 	}

// 	close(odd_ch)
// 	close(even_ch)
// 	close(flag)

// 	wg.Wait()

// }

// func print_even(even_ch chan int, flag chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for val := range even_ch {
// 		fmt.Println("even", val)
// 		flag <- true
// 	}

// }

// func print_odd(odd_ch chan int, flag chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for val := range odd_ch {
// 		fmt.Println("odd", val)
// 		flag <- true
// 	}

// }

//tests
//tests

// tests

// tests

//tests

// func main() {
// 	//declared the var channel in stack memory, not allocated memory in the heap, that's why its value is nil
// 	var ch chan int
// 	fmt.Println("start Main method")
// 	// unbuffered channel ready to receicve the value and memory allocated in heap memory
// 	//ch := make(chan int)
// 	fmt.Println(ch)
// 	go print(ch)

// 	ch <- 5
// 	fmt.Println("end Main method")

// }

// func print(ch chan int) {
// 	fmt.Println(<-ch)
//}
//
///
//
//
//
/*
















 */

//  1. It has to be at least 16 characters long.
// 2. The password cannot contain the word "password". This rule is not case-sensitive.
// 3. The same character cannot be used more than 4 times. This rule is case-sensitive, "a" and "A" are different characters.
// Write a function that takes in a password and returns a collection of any rule numbers that are not met.

// password_1 = "Strongpwd9999#abc"             ==> []
// password_2 = "Less10#"                       ==> [1]
// password_3 = "Password@"                     ==> [1,2]
// password_4 = "#PassWord0aaaa11111112222223x"     ==> [2,3]

func main() {
	//str := "#PassWord0aaaa11111112222223x"

	str := "Password@"

	fmt.Println(ValidPassword(str))
}

func ValidPassword(str string) []int {

	var res []int

	if len(str) < 16 {
		res = append(res, 1)
	}

	str = strings.ToLower(str)

	if strings.Contains(str, "password") {

		res = append(res, 2)
	}

	map1 := make(map[rune]int)

	for i := 0; i < len(str); i++ {
		if _, ok := map1[rune(str[i])]; !ok {
			map1[rune(str[i])] = 1
		} else if map1[rune(str[i])] >= 4 {
			res = append(res, 3)
			return res
		} else {
			count := map1[rune(str[i])]
			map1[rune(str[i])] = count + 1
		}
	}

	return res
}
