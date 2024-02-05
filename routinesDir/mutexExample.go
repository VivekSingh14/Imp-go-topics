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
	mutex.Lock()

	Balance += value
	mutex.Unlock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, Balance)

}

func Withdraw(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()

	Balance -= value
	mutex.Unlock()
	fmt.Printf("Withdrwaing %d from account with balance: %d\n", value, Balance)

}
