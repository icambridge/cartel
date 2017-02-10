package cartel

import (
	"runtime"
	"testing"
	"time"
)

func Test_Create_Starts_Correct_Number_Of_Goroutines(t *testing.T) {

	startNumber := runtime.NumGoroutine()
	numberOfWorkers := 5

	p := NewPool(PoolOptions{Size: numberOfWorkers})
	defer p.End()
	actualNumber := runtime.NumGoroutine()
	expectedNumber := startNumber + numberOfWorkers

	if actualNumber != expectedNumber {
		t.Errorf("expected %v goroutines but got %v goroutines", expectedNumber, actualNumber)

	}
}

func Test_End_Kills_Goroutines(t *testing.T) {

	startNumber := runtime.NumGoroutine()
	numberOfWorkers := 5

	p := NewPool(PoolOptions{Size: numberOfWorkers})
	p.End()
	actualNumber := runtime.NumGoroutine()

	if actualNumber != startNumber {
		t.Errorf("expected %v goroutines but got %v goroutines", startNumber, actualNumber)
	}

}

func Test_Returns_Output(t *testing.T) {

	numberOfWorkers := 5

	p := NewPool(PoolOptions{Size: numberOfWorkers})

	task := TestTask{"Iain"}

	p.Do(task)
	p.End()

	values := p.GetOutput()

	if expected, actual := 1, len(values);expected != actual {
		t.Errorf("expected %v  but got %v ", expected, actual)
	}
	value := values[0]
	if expected, actual := "Iain", value; expected != actual {
		t.Errorf("expected %v  but got %v ", expected, actual)
	}
}

func Test_GetOutput(t *testing.T) {

	numberOfWorkers := 5

	p := NewPool(PoolOptions{Size: numberOfWorkers})

	task := TestTask{"Iain"}

	p.Do(task)
	p.End()

	outputs := p.GetOutput()
	value := outputs[0]
	if expected, actual := "Iain", value; expected != actual {
		t.Errorf("expected %v  but got %v ", expected, actual)
	}
}

type TestReceiver struct {
	Output string
}

type TestTask struct {
	Name string
}

func (tt TestTask) Execute() interface{} {
	return tt.Name
}

type TestTimeTask struct {
	Name string
}

func (ttt TestTimeTask) Execute() interface{} {
	return time.Now()
}
