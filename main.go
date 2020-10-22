package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var waitGroup sync.WaitGroup

	//Check the signal for daemon usage
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		s := <-sigs
		if s == syscall.SIGURG {
			fmt.Println("received sigur")
		} else {
			log.Printf("RECEIVED SIGNAL: %s\n", s)
			os.Exit(1)
		}
	}()

	//flags to be used
	//Example : toolname -f to flag1
	//By default, toolanme -h will print list of flags with description

	flag1 := flag.String("f", "default", "flag description for -h")
	flag2 := flag.String("t", "default", "flag description for -h")

	flag.Parse()

	// OPTIONAL - No flag catcher
	if (*flag1 == "") && (*flag2 == "") {
		fmt.Printf("Error, need at least an argument")
		os.Exit(0)
	}

	//goRoutines to launch. Replace numberOfGoroutines with an int or a var that represent the number of goroutines you want to asynchronously launch
	waitGroup.Add(numberOfGoroutines)

	waitGroup.wait
	os.Exit(0)
}
