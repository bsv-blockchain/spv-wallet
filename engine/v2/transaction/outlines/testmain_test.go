package outlines_test

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	code := m.Run()

	// Wait briefly for cleanup
	time.Sleep(100 * time.Millisecond)
	goroutines := runtime.NumGoroutine()

	// If tests passed and there are lingering goroutines from external libraries
	// exit immediately to prevent CI timeout
	if code == 0 && goroutines > 5 {
		fmt.Printf("[TestMain:outlines] Tests passed but %d goroutines still running, exiting cleanly\n", goroutines)
		os.Exit(0)
	}

	os.Exit(code)
}
