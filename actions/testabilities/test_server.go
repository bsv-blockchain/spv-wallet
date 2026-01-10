package testabilities

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type testServer struct {
	handlers *gin.Engine
}

func (t *testServer) RoundTrip(request *http.Request) (*http.Response, error) {
	r := httptest.NewRecorder()
	t.handlers.ServeHTTP(r, request)

	// Give background goroutines (like database/sql's context monitoring) time to
	// finish reading from gin's context before it gets returned to the pool and
	// reused by the next request. This helps prevent races on gin.Context.Request.
	// Note: We can't use a mutex here because the server may make reentrant HTTP
	// calls (e.g., paymail capabilities) that go through this same RoundTrip.
	runtime.Gosched()
	time.Sleep(20 * time.Millisecond)

	return r.Result(), nil
}
