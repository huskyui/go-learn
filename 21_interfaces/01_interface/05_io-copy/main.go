package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "qdajshadbasdhgasjhd"
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
	res, _ := http.Get("https://www.bilibili.com")
	io.Copy(os.Stdout, res.Body)
	fmt.Printf("%T", res.Body.Close())

}
