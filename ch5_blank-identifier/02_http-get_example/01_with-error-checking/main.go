package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	page, error := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s", page)

}
