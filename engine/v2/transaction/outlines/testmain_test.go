package outlines_test

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Printf("[TestMain:outlines] Starting tests - goroutines: %d\n", runtime.NumGoroutine())

	code := m.Run()

	fmt.Printf("[TestMain:outlines] Tests completed with code %d - goroutines: %d\n", code, runtime.NumGoroutine())

	// Wait briefly for cleanup
	time.Sleep(100 * time.Millisecond)
	goroutines := runtime.NumGoroutine()
	fmt.Printf("[TestMain:outlines] After 100ms - goroutines: %d\n", goroutines)

	// If tests passed and there are lingering goroutines from external libraries
	// exit immediately to prevent CI timeout
	if code == 0 && goroutines > 5 {
		fmt.Printf("[TestMain:outlines] Tests passed but %d goroutines still running (from external libraries), exiting cleanly\n", goroutines)
		os.Exit(0)
	}

	os.Exit(code)
}
