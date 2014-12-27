package cartel

import (
	"runtime"
	"testing"
)

func Test_Create_Starts_Correct_Number_Of_Goroutines(t *testing.T) {
	
	startNumber := runtime.NumGoroutine()
	numberOfWorkers := 5
	
	NewPool(numberOfWorkers)
	
	actualNumber := runtime.NumGoroutine()
	expectedNumber := startNumber+numberOfWorkers
	
	if actualNumber != expectedNumber {
		t.Errorf("expected %v goroutines but got %v goroutines", expectedNumber, actualNumber)
		
	}
}

func Test_End_Kills_Goroutines(t *testing.T) {

	startNumber := runtime.NumGoroutine()
	numberOfWorkers := 5

	p := NewPool(numberOfWorkers)
	p.End()
	actualNumber := runtime.NumGoroutine()

	if actualNumber != startNumber {
		t.Errorf("expected %v goroutines but got %v goroutines", startNumber, actualNumber)
	}

}

func Test_Returns_Output(t *testing.T) {

	p := NewPool(1)
	
	task := TestTask{"Iain"}
	
	p.Do(task)
	p.End()
	
	value := <-p.Output
	
	if expected, actual := "Iain", value.Value(); expected != actual {
		t.Errorf("expected %v  but got %v ", expected, actual)
	}
}


func Test_GetOutput(t *testing.T) {
	
	p := NewPool(1)
	
	task := TestTask{"Iain"}

	p.Do(task)
	p.End()

	outputs := p.GetOutput()
	value := outputs[0]
	if expected, actual := "Iain", value.Value(); expected != actual {
		t.Errorf("expected %v  but got %v ", expected, actual)
	}
}

type TestReceiver struct {
	Output string
}

type TestTask struct {
	Name string
}

func (tt TestTask) Execute() OutputValue {
	return TestOutput{tt.Name}
}

type TestOutput struct {
	Name string
}

func (to TestOutput) Value() interface {} {
	return to.Name
}
