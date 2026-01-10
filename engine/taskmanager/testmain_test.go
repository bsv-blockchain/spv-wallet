package taskmanager

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Printf("[TestMain] Starting tests - goroutines: %d\n", runtime.NumGoroutine())

	code := m.Run()

	fmt.Printf("[TestMain] Tests completed with code %d - goroutines: %d\n", code, runtime.NumGoroutine())

	// Wait briefly and check goroutine count again
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("[TestMain] After 100ms - goroutines: %d\n", runtime.NumGoroutine())

	// Print goroutine stacktraces if there are more than expected (typically 2-3 for Go runtime)
	if runtime.NumGoroutine() > 5 {
		fmt.Println("[TestMain] WARNING: More goroutines than expected, dumping stacks:")
		buf := make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		fmt.Printf("%s\n", buf[:n])
	}

	os.Exit(code)
}
