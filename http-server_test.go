package cheatsheet

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
}

func TestHTTPServer(t *testing.T) {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
		return
	}
}
