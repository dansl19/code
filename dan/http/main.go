package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if resp != nil {
		fmt.Println("this is a response", resp)
	}
	if err != nil {
		fmt.Println("this is an errro", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999) // here we create "bs" var from type slice with lenght of 99999 in order to be able to stole whole response in this slice
	resp.Body.Read(bs)        //here we take the response "resp" call on this response function "Body" and then function "Read" and store everyhting in "bs" var
	fmt.Println(string(bs))   // here we print content of "bs" var
}
