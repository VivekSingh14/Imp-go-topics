package routinesdir

func Ping(pings chan<- string, msg string) {
	pings <- msg
}

func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
