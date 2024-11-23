package cheatsheet

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	var pongWait = 10 * time.Second
	fmt.Println("Pong Wait = ", pongWait)

	var pingInterval = (pongWait * 9) / 10
	fmt.Println("Ping Interval = ", pingInterval)
}
