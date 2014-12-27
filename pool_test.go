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
