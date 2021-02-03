package main

import (
	"log"
	"microservice-with-go/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "Allen Micro-Service:", log.LstdFlags)
	l.Println("Start Service...")

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":8080", sm)
}
