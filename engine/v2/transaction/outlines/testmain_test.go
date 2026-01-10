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

	// Wait briefly and check goroutine count again
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("[TestMain:outlines] After 100ms - goroutines: %d\n", runtime.NumGoroutine())

	// Give more time for cleanup
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[TestMain:outlines] After 600ms - goroutines: %d\n", runtime.NumGoroutine())

	// Print goroutine stacktraces if there are more than expected
	if runtime.NumGoroutine() > 5 {
		fmt.Println("[TestMain:outlines] WARNING: More goroutines than expected, dumping stacks:")
		buf := make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		fmt.Printf("%s\n", buf[:n])
	}

	os.Exit(code)
}
