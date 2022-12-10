package routinesdir

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func GetData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	response.Body.Close()

}
