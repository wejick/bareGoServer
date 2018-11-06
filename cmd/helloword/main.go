package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/bareGoServer/internal/helloWorld"
)

func main() {
	router := httprouter.New()

	hello := helloWorld.NewHelloWorld("hallo")

	// swagger:route GET /helloword/{name} hellowordWithName
	//
	// Responses:
	//  200: HellowordResponse
	router.GET("/helloword/:name", hello.HandlerHelloWorldWithName)

	log.Fatal(http.ListenAndServe(":8080", router))
}
