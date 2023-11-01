package main

import (
	"testing"
)

func TestPerson_Greet(t *testing.T) {
	person := Person{Name: "John Doe", Age: 30}

	expectedOutput := "Hello, my name is John Doe\n"
	if testing.Verbose() {
		expectedOutput += "Hello, my name is John Doe\n"
	}

	if testing.Verbose() {
		t.Log("Expected output:", expectedOutput)
	}

	if testing.Verbose() {
		t.Log("Actual output:")
	}

	capturedOutput := captureOutput(func() {
		person.Greet()
	})

	if capturedOutput != expectedOutput {
		t.Errorf("Expected greeting '%s', but got '%s'", expectedOutput, capturedOutput)
	}
}

func captureOutput(f func()) string {
	var capturedOutput string
	outputCapture := func(args ...interface{}) {
		capturedOutput = fmt.Sprint(args...)
	}

	oldOutput := os.Stdout
	pos.Stdout = outputCapture

	defer func() { os.Stdout = oldOutput }()

	f()

	return capturedOutput
}