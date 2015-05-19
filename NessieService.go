package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	)

func main() {
	key := "ba5d1d41f788971f4816f3d7a0e644ff"
	url := fmt.Sprintf("http://api.reimaginebanking.com/atms?key=%s",key)
    res, err := http.Get(url)
    errCheck(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    errCheck(err)
    fmt.Printf("%s\n", string(body))
}
func errCheck(e error) {
    if e != nil {
        panic(e)
    }
}

