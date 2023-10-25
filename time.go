package main

import (
	"fmt"
	"time"
)

func main() {
	var pongWait = 10 * time.Second
	fmt.Println("Pong Wait = ", pongWait)

	var pingInterval = (pongWait * 9) / 10
	fmt.Println("Ping Interval = ", pingInterval)
}
