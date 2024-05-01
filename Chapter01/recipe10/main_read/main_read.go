package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	cmd := "ping"
	timeout := 2 * time.Second

	// The command line tool
	// "ping" is executed for
	// 2 seconds
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	proc := exec.CommandContext(ctx, cmd, "example.com")

	// The process output is obtained
	// in form of io.ReadCloser. The underlying
	// implementation use the os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Start the process
	err := proc.Start()
	if err != nil {
		log.Panic(err)
	}

	// For more comfortable reading the
	// bufio.Scanner is used.
	// The read call is blocking.
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
