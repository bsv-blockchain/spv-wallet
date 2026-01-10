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

	// Use a channel to run tests with a timeout
	done := make(chan int, 1)
	go func() {
		done <- m.Run()
	}()

	// Wait for tests to complete with a 2-minute timeout
	select {
	case code := <-done:
		fmt.Printf("[TestMain:outlines] Tests completed with code %d - goroutines: %d\n", code, runtime.NumGoroutine())

		// Wait briefly for cleanup
		time.Sleep(100 * time.Millisecond)
		goroutines := runtime.NumGoroutine()
		fmt.Printf("[TestMain:outlines] After 100ms - goroutines: %d\n", goroutines)

		// Force exit regardless of goroutine count for clean CI behavior
		if code == 0 {
			fmt.Printf("[TestMain:outlines] Tests passed, exiting cleanly\n")
		}
		os.Exit(code)

	case <-time.After(2 * time.Minute):
		fmt.Printf("[TestMain:outlines] Tests timed out after 2 minutes - goroutines: %d\n", runtime.NumGoroutine())
		os.Exit(1)
	}
}
