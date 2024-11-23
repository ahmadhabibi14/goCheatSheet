package cheatsheet

import (
	"fmt"
	"os/exec"
	"runtime"
	"testing"
)

func TestExecV2(t *testing.T) {
	var output []byte
	var err error
	if runtime.GOOS == `windows` {
		output, err = exec.Command("cmd", "/c", "git config user.name").Output()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		output, err = exec.Command("bash", "-c", "git config user.name").Output()
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Operating System: ", runtime.GOOS)
	fmt.Printf("-> Git username:\n%s\n", string(output))
}
