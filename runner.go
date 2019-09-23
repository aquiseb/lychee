package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type cmd struct {
	Module string
	Cmd    string
	Args   []string
	Dir    string
	Err    error
}

func runner() {

	torun := []cmd{
		cmd{
			Module: "Hello Microservice",
			Cmd:    "go",
			Args:   []string{"run", "main.go"},
			Dir:    "./micro-hello",
		},
		cmd{
			Module: "Post Microservice",
			Cmd:    "go",
			Args:   []string{"run", "main.go"},
			Dir:    "./micro-post",
		},
	}

	var wg sync.WaitGroup // use a waitgroup to ensure all concurrent jobs are running
	wg.Add(len(torun))

	out := make(chan cmd)     // a channel to output cmd status
	succ := make(chan string) // a channel to output cmd status

	go func() {
		wg.Wait() //wait for the group to finish
		cm := exec.Command("go", "run", "micro-federation/main.go")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		cm.Stdout = mw
		cm.Stderr = mw
		cm.Run()
		close(out) //  then close the signal channel
	}()

	// start the commands
	for _, c := range torun {
		// go runCmd(c, out, &wg)
		go runCmdAndWaitForSomeOutput(c, out, succ, &wg)
	}

	// loop over the chan to collect errors
	// it ends when wg.Wait unfreeze and closes out
	for c := range out {
		if c.Err != nil {
			log.Fatalf("%v %v has failed with %v", c.Cmd, c.Args, c.Err)
		}
	}

	// here all commands started you can proceed further to run the last command
	fmt.Println("all done")
	os.Exit(0)
}

func runCmd(o cmd, out chan cmd, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command(o.Cmd, o.Args...)
	if o.Dir != "" {
		cmd.Dir = o.Dir
	}

	if err := cmd.Start(); err != nil {
		o.Err = err // save err
		out <- o    // signal completion error
		return      // return to unfreeze the waitgroup wg
	}
	go cmd.Wait() // dont wait for command completion,
	// consider its done once the program started with success.

	// out <- o // useless as main look ups only for error
}

func runCmdAndWaitForSomeOutput(o cmd, out chan cmd, succ chan string, wg *sync.WaitGroup) {

	cmd := exec.Command(o.Cmd, o.Args...)
	if o.Dir != "" {
		cmd.Dir = o.Dir
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		o.Err = err // save err
		out <- o    // signal completion
		return      // return to unfreeze the waitgroup wg
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		o.Err = err
		out <- o
		return
	}

	if err := cmd.Start(); err != nil {
		o.Err = err
		out <- o
		return
	}

	go cmd.Wait() // dont wait for command completion

	outScan := make(chan error) // to signal errors detected on the fd
	var wg2 sync.WaitGroup
	wg2.Add(2) // the number of fds being watched

	go func() {
		defer wg.Done()
		defer wg2.Done()
		sc := bufio.NewScanner(stdout)
		for sc.Scan() {
			line := sc.Text()
			fmt.Println(line)
			if strings.Contains(line, "FEDERATION_SIGNAL_OK") { // the OK marker
				return // quit asap to unfreeze wg2
			} else if strings.Contains(line, "not known") { // the nOK marker, if any...
				outScan <- fmt.Errorf("%v", line)
				return // quit  to unfreeze wg2
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer wg2.Done()
		sc := bufio.NewScanner(stderr)
		for sc.Scan() {
			line := sc.Text()
			fmt.Println(line)
			if strings.Contains(line, "FEDERATION_SIGNAL_OK") { // the OK marker
				return // quit asap to unfreeze wg2
			} else if strings.Contains(line, "not known") { // the nOK marker, if any...
				outScan <- fmt.Errorf("%v", line) // signal error
				return                            // quit to unfreeze wg2
			}
		}
	}()

	// go func() {
	wg2.Wait() // consider that if the program does not output anything,
	// or never prints ok/nok, this will block forever
	close(outScan) // close the chan so the next loop is finite
	// }()

	// - simple timeout less loop
	for err := range outScan {
		if err != nil {
			o.Err = err // save the execution error
			out <- o    // signal the cmd
			return      // qui to unfreeze the wait group wg2
		}
	}

	// - more complex version with timeout
	// timeout := time.After(time.Second * 3)
	// for {
	// 	select {
	// 	case err, ok := <-outScan:
	// 		if !ok { // if !ok, outScan is closed and we should quit the loop
	// 			return
	// 		}
	// 		if err != nil {
	// 			o.Err = err // save the execution error
	// 			out <- o    // signal the cmd
	// 			return      // quit to unfreeze the wait group wg2
	// 		}
	// 	case <-timeout:
	// 		o.Err = fmt.Errorf("timed out...%v", timeout) // save the execution error
	// 		out <- o                                      // signal the cmd
	// 		return                                        // quit to unfreeze the wait group wg2
	// 	}
	// }

	// exit and unfreeze the wait group wg
}
