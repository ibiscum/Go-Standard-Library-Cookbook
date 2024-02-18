package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	pReader, pWriter := io.Pipe()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("C:/daten/msys64/usr/bin/echo.exe", "Hello Go!\nThis is an example")
	} else {
		cmd = exec.Command("/usr/bin/echo", "Hello Go!\nThis is an example")
	}
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
