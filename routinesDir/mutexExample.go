package routinesdir

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	Balance int
)

func init() {
	Balance = 1000
}

func Deposit(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	// mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, Balance)
	Balance += value
	// mutex.Unlock()

}

func Withdraw(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	// mutex.Lock()
	fmt.Printf("Withdrwaing %d from account with balance: %d\n", value, Balance)
	Balance -= value
	// mutex.Unlock()

}
