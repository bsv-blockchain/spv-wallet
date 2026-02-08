package taskmanager

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTaskManager_Close_WaitsForCronJobs(t *testing.T) {
	if raceEnabled {
		t.Skip("skipping due to internal data race in taskq consumer (vmihailenco/taskq/v3 Consumer.Add vs Consumer.worker)")
	}

	ctx := context.Background()
	tm, err := NewTaskManager(ctx)
	require.NoError(t, err)
	require.NotNil(t, tm)

	// Track job execution with channels for synchronization
	jobStarted := make(chan struct{})
	jobCompleted := make(chan struct{})

	// Add a cron job that signals when it starts and completes
	err = tm.(*TaskManager).CronJobsInit(CronJobs{
		"slow_job_close_test": {
			Period: 100 * time.Millisecond,
			Handler: func(_ context.Context) error {
				close(jobStarted) // Signal job started
				time.Sleep(200 * time.Millisecond)
				close(jobCompleted) // Signal job completed
				return nil
			},
		},
	})
	require.NoError(t, err)

	// Wait for the cron job to actually start running
	select {
	case <-jobStarted:
		// Job has started, now we can test Close
	case <-time.After(2 * time.Second):
		t.Fatal("cron job did not start within timeout")
	}

	// Close should wait for the running job to complete
	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = tm.Close(closeCtx)
	require.NoError(t, err)

	// Verify job completed - check immediately since Close should have waited
	select {
	case <-jobCompleted:
		// Job completed as expected
	default:
		t.Fatal("cron job should have completed before Close returned")
	}
}

func TestTaskManager_Close_RespectsContextTimeout(t *testing.T) {
	if raceEnabled {
		t.Skip("skipping due to internal data race in taskq consumer (vmihailenco/taskq/v3 Consumer.Add vs Consumer.worker)")
	}

	ctx := context.Background()
	tm, err := NewTaskManager(ctx)
	require.NoError(t, err)
	require.NotNil(t, tm)

	// Close with a very short timeout - should not hang
	closeCtx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	err = tm.Close(closeCtx)
	elapsed := time.Since(start)

	// Close should return quickly, not hang forever
	require.NoError(t, err)
	require.Less(t, elapsed, 500*time.Millisecond, "Close should return promptly, not hang")
}

func TestTaskManager_Close_Idempotent(t *testing.T) {
	ctx := context.Background()
	tm, err := NewTaskManager(ctx)
	require.NoError(t, err)
	require.NotNil(t, tm)

	// First close should succeed
	err = tm.Close(context.Background())
	require.NoError(t, err)

	// Second close should also succeed (idempotent)
	err = tm.Close(context.Background())
	require.NoError(t, err)
}

func TestTaskManager_Close_WithNilContext(t *testing.T) {
	ctx := context.Background()
	tm, err := NewTaskManager(ctx)
	require.NoError(t, err)
	require.NotNil(t, tm)

	// Close with background context should work
	err = tm.Close(context.Background())
	require.NoError(t, err)
}
