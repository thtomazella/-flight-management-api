package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Teste de abertura")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
