package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from golang http srver")
	w.Write([]byte("ok"))
}

func sayHelloName2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from golang http srver")
	w.Write([]byte("ok"))
}

func sayHelloName3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from golang http srver")
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", sayHelloName1)
	http.HandleFunc("/", sayHelloName2)
	http.HandleFunc("/", sayHelloName3)
	err := http.ListenAndServe("0.0.0.0:9898", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
