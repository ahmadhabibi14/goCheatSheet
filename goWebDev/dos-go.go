package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	json "github.com/goccy/go-json"
)

type (
	Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func httpRequest(payload []byte) {
	for {
		resp, err := http.Post(
			"http://localhost:1234/guest/login",
			"application/json",
			bytes.NewBuffer(payload),
		)
		if err != nil {
			log.Fatalln("Error: ", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Login successful !")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	payload := Payload{
		Email:    "tes1234@proton.me",
		Password: "tes1234567890",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	go httpRequest(jsonPayload)
	select {}
}
