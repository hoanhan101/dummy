package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
)

func main() {
	resp, err := http.Get("http://github.com/hoanhan101")
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
	defer resp.Body.Close()
	fmt.Printf("%#v\n", resp.Body)

    bytes_buffer , err := ioutil.ReadAll(resp.Body)
    fmt.Printf("%T\n", string(bytes_buffer))
}
