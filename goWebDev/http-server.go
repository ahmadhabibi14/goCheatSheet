package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	/*
		This is a Go function that takes ResponseWriter and Request as an input and writes
		Hello World! on an HTTP response stream.
	*/
	fmt.Fprintf(w, "Hello World !")
}

/*
Next, we declared the main() method from where the program execution begins, as this
method does a lot of things. Let’s understand it line by line:
*/
func main() {
	http.HandleFunc("/", helloWorld)
	/*
		Here, we are registering the helloWorld function with
		the / URL pattern using HandleFunc of the net/http package, which means
		helloWorld gets executed, passing (http.ResponseWriter, *http.Request)
		as a parameter to it whenever we access the HTTP URL with pattern /.
	*/
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
		return
	}
}

/*
	HOW IT WORKS ...

	Once we run the program, an HTTP server will start locally listening on port 8080.
	Opening http://localhost:8080 in a browser will display Hello World! from the server.
*/
