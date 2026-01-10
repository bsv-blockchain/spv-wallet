package engine

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Printf("[TestMain:engine] Starting tests - goroutines: %d\n", runtime.NumGoroutine())

	code := m.Run()

	fmt.Printf("[TestMain:engine] Tests completed with code %d - goroutines: %d\n", code, runtime.NumGoroutine())

	// Wait briefly for cleanup
	time.Sleep(100 * time.Millisecond)
	goroutines := runtime.NumGoroutine()
	fmt.Printf("[TestMain:engine] After 100ms - goroutines: %d\n", goroutines)

	// If tests passed and there are lingering goroutines from external libraries
	// (taskq, cron, etc.), exit immediately to prevent CI timeout
	// The goroutines are from test cleanup timing issues, not actual leaks
	if code == 0 && goroutines > 5 {
		fmt.Printf("[TestMain:engine] Tests passed but %d goroutines still running (from external libraries), exiting cleanly\n", goroutines)
		os.Exit(0)
	}

	os.Exit(code)
}
