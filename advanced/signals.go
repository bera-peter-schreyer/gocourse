package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signals() {
	pid := os.Getpid()
	fmt.Printf("Process ID: %d\n", pid)
	sigs := make(chan os.Signal, 1)
	go func() {
		sig := <-sigs
		fmt.Print("Received signal: ", sig, "\n")
		switch sig {
		case syscall.SIGINT:
			println("Received SIGINT, exiting...")
			os.Exit(0)
		case syscall.SIGTERM:
			println("Received SIGTERM, exiting...")
			os.Exit(0)
		default:
			println("Received signal:", sig)
		}		
	}()
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGHUP)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)
	signal.Notify(sigs, syscall.SIGPIPE, syscall.SIGCHLD)
	signal.Notify(sigs, syscall.SIGCONT, syscall.SIGSTOP)
	signal.Notify(sigs, syscall.SIGTSTP, syscall.SIGTTIN)
	for {
		time.Sleep(1 * time.Second)
	}
}