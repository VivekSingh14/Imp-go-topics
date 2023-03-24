package routinesdir

import "fmt"

func Ting(pings <-chan string) {

	msg := <-pings

	fmt.Println(msg, " from Ting")

}

func Tong(pings <-chan string) {

	msg := <-pings

	fmt.Println(msg, " from Tong")
}
