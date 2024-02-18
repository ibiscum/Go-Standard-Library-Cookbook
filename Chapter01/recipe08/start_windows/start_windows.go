//go:build windows

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	prc := exec.Command("C:/daten/msys64/usr/bin/ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	err = prc.Wait()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
