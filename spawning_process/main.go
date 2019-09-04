package main

import (
	"fmt"
	"os/exec"
)

func main() {
	defer func() {
		if r := recover(); r!= nil {
			fmt.Println("recovered from", r)
		}
	}()
	dateCmd := exec.Command("dat")
	dateCmdResult, err := dateCmd.Output()
	if err != nil {
		panic("panic executing program")
	}
	fmt.Println(string(dateCmdResult))
}
