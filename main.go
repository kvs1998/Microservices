package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	//servemux is an handler
	//which handler registered against it get called
	//and then just call servehttp func of the handler based on
	// the request
	sm := http.NewServeMux() //serve mux created
	ss.Handle("/", hh)       //register path to hh-hello handler
	http.ListenAndServe(":9090", nil)
}
