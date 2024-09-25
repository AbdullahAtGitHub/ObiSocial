package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server"))
}

func main() {
	s := &server{addr: ":8080"}
	log.Printf("Server is started at port %v \n", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, s))
}
