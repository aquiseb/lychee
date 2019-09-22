package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var modulesToRun = []string{"micro-post", "micro-hello"}

func main() {
	// Send multiple values to chan
	// https://stackoverflow.com/a/50857250/9077800
	c := make(chan func() (string, error))

	go runModule([]string{"go", "run", "micro-post"}, c)  // PROCESS A
	go runModule([]string{"go", "run", "micro-hello"}, c) // PROCESS B

	modulesRunning := []string{}
	for {
		msg, err := (<-c)()
		if err != nil {
			log.Fatalln(err)
		}

		if strings.HasPrefix(msg, "micro-") && err == nil {
			modulesRunning = append(modulesRunning, msg)
			if CompareUnorderedSlices(modulesToRun, modulesRunning) {
				go runModule([]string{"go", "run", "micro-federation"}, c) // PROCESS C
			}
		}
	}

}

func runModule(commandArgs []string, o chan func() (string, error)) {
	cmd := exec.Command(commandArgs[0], commandArgs[1], commandArgs[2]+"/main.go")

	// Less verbose solution to stream output with io?
	// var stdBuffer bytes.Buffer
	// mw := io.MultiWriter(os.Stdout, &stdBuffer)
	// cmd.Stdout = mw
	// cmd.Stderr = mw

	c := make(chan struct{})
	wg.Add(1)

	// Stream command output
	// https://stackoverflow.com/a/38870609/9077800
	go func(cmd *exec.Cmd, c chan struct{}) {
		defer wg.Done()
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			close(o)
			panic(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			close(o)
			panic(err)
		}

		<-c
		outScanner := bufio.NewScanner(stdout)
		for outScanner.Scan() {
			m := outScanner.Text()
			fmt.Println(commandArgs[2]+":", m)
			o <- (func() (string, error) { return commandArgs[2], nil })
		}

		errScanner := bufio.NewScanner(stderr)
		for errScanner.Scan() {
			m := errScanner.Text()
			fmt.Println(commandArgs[2]+":", m)
			o <- (func() (string, error) { return "bad", nil })
		}
	}(cmd, c)

	c <- struct{}{}
	cmd.Start()

	wg.Wait()
	close(o)
}

// CompareUnorderedSlices orders slices before comparing them
func CompareUnorderedSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	return reflect.DeepEqual(a, b)
}
