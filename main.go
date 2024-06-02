package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Hello world!")
}

func main(){
    http.HandleFunc("/", helloWorldPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}