package main

import (
	"io"
	"os/exec"
)

func osProcesses() {

	pr, pw := io.Pipe()
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr
	go func() {
		defer pw.Close()
		pw.Write([]byte("foo\nbar\nbaz\n"))
	}()
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	println(string(output))

	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(output))

	// cmd := exec.Command("sleep", "60")
	// // Start the command
	// err := cmd.Start()
	// if err != nil {
	// 	panic(err)
	// }
	// time.Sleep(time.Second * 2)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Command killed successfully")
	// // Wait for the command to finish
	// err = cmd.Wait()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Command finished successfully")

	// cmd := exec.Command("grep", "foo")
	// // Set the input for the command
	// cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(output))
	
	// cmd := exec.Command("echo", "Hello, World!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(output))

}