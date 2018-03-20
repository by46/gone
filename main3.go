package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("main3.exe", `-xarg:sleep 1`)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s", out.String())
}
