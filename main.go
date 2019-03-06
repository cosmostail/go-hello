package main

import (
	"fmt"
	"log"
	"net/http"
	"./hello"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hello.SayHello() + "Scott")
}

func main()  {
	//fmt.Println("hello world")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}