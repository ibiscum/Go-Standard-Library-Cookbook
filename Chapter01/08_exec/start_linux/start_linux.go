//go:build linux

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	err = prc.Wait()
	if err != nil {
		log.Panic(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
