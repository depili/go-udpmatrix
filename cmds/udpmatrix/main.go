package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

var options struct {
	UdpListen  string `short:"l" long:"listen" description:"UDP port to listen to" default:":4242"`
	Rows       int    `short:"r" long:"rows" description:"Rows per led module" default:"32"`
	Chain      int    `short:"c" long:"chain" description:"Chained panels" default:"6"`
	Parallel   int    `short:"p" long:"paraller" description:"Parallel chains" default:"3"`
	Brightness int    `short:"b" long:"brightness" description:"Brightness 0-100" default:"100"`
}

var parser = flags.NewParser(&options, flags.Default)

func main() {
	// Parse flags
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	fmt.Printf("Starting\n")
	c := initMatrix()
	go runListener(c)
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
