package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	expectedOutput :=
		fmt.Sprintln("Hello, universe!") +
			fmt.Sprintln("Hello, cosmos!") +
			fmt.Sprintln("Hello, world!")

	out, _ := io.ReadAll(r)
	actualOutput := string(out)

	os.Stdout = stdOut

	if actualOutput != expectedOutput {
		t.Errorf("\nExpected: \n%s\nGot: \n%s\n", expectedOutput, actualOutput)
	}

}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "hello, there"
	printMessage()

	_ = w.Close()

	out, _ := io.ReadAll(r)
	actual := string(out)
	expected := fmt.Sprintln(msg)

	os.Stdout = stdOut

	if actual != expected {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, actual)
	}

	msg = ""

}

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("hello, there", &wg)
	wg.Wait()

	if msg != "hello, there" {
		t.Errorf("\nExpected: hello, there\nGot: %s", msg)
	}
}
