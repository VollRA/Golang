package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body := bytes.NewBuffer([]byte("Ololololo"))
	r, _ := http.Post("http://localhost:8081/read", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}
