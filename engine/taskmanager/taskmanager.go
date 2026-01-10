/*
Package taskmanager is the task/job management service layer for concurrent and asynchronous tasks with cron scheduling.
*/
package taskmanager

import (
	"context"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/vmihailenco/taskq/v3"

	"github.com/bsv-blockchain/spv-wallet/engine/logging"
)

type (

	// TaskManager implements the TaskEngine interface
	TaskManager struct {
		options *options
	}

	options struct {
		cronService *cron.Cron      // Internal cron job client
		cronMu      sync.Mutex      // Mutex for cronService operations to prevent race conditions
		logger      *zerolog.Logger // Internal logging
		taskq       *taskqOptions   // All configuration and options for using TaskQ
	}

	// taskqOptions holds all the configuration for the TaskQ engine
	taskqOptions struct {
		config   *taskq.QueueOptions    // Configuration for the TaskQ engine
		queue    taskq.Queue            // Queue for TaskQ
		consumer taskq.QueueConsumer    // Consumer for TaskQ (Redis only)
		tasks    map[string]*taskq.Task // Registered tasks
		queueMu  sync.Mutex             // Mutex for queue operations to prevent race conditions
		tasksMu  sync.RWMutex           // Mutex for tasks map operations
	}
)

// NewTaskManager creates a new client for all TaskManager functionality
// If no options are given, it will use local memory for the queue.
func NewTaskManager(ctx context.Context, opts ...Options) (TaskEngine, error) {
	// Create a new tm with defaults
	tm := &TaskManager{options: &options{
		taskq: &taskqOptions{
			tasks:  make(map[string]*taskq.Task),
			config: DefaultTaskQConfig("taskq"),
		},
	}}

	// Overwrite defaults with any set by user
	for _, opt := range opts {
		opt(tm.options)
	}

	if tm.options.logger == nil {
		tm.options.logger = logging.GetDefaultLogger()
	}

	if err := tm.loadTaskQ(ctx); err != nil {
		return nil, err
	}

	tm.ResetCron()

	return tm, nil
}

// Close the client and any open connections
func (tm *TaskManager) Close(ctx context.Context) error {
	if tm != nil && tm.options != nil {

		// Stop the cron scheduler and wait for running jobs to complete
		// cron.Stop() returns a context that signals when all jobs are done
		tm.options.cronMu.Lock()
		var cronCtx context.Context
		if tm.options.cronService != nil {
			cronCtx = tm.options.cronService.Stop()
			tm.options.cronService = nil
		}
		tm.options.cronMu.Unlock()

		// Wait for cron jobs to complete (if any were running) with timeout
		if cronCtx != nil {
			select {
			case <-cronCtx.Done():
				// Cron stopped cleanly, all jobs completed
			case <-time.After(200 * time.Millisecond):
				// Timeout waiting for cron jobs, proceed with cleanup
				tm.options.logger.Warn().Msg("timeout waiting for cron jobs to complete")
			case <-ctx.Done():
				// Parent context canceled, proceed with cleanup
			}
		}

		// Stop the consumer before closing the queue (Redis only)
		// Use StopTimeout to avoid the default 30 second timeout
		// Use 1s timeout to allow internal goroutines to complete cleanly
		if tm.options.taskq.consumer != nil {
			if err := tm.options.taskq.consumer.StopTimeout(1 * time.Second); err != nil {
				tm.options.logger.Warn().Err(err).Msg("error stopping taskq consumer")
			}
		}

		// Close the taskq queue with short timeout (protected by mutex to prevent race with Add operations)
		// Use CloseTimeout to avoid the default 30 second timeout
		// Use 1s timeout to allow internal goroutines to complete cleanly
		tm.options.taskq.queueMu.Lock()
		if tm.options.taskq.queue != nil {
			queue := tm.options.taskq.queue
			tm.options.taskq.queue = nil
			tm.options.taskq.queueMu.Unlock()

			if err := queue.CloseTimeout(1 * time.Second); err != nil {
				tm.options.logger.Warn().Err(err).Msg("error closing taskq queue")
			}
		} else {
			tm.options.taskq.queueMu.Unlock()
		}

		// Empty all values and reset
		tm.options.taskq.config = nil
	}

	return nil
}

// ResetCron will reset the cron scheduler and all loaded tasks
func (tm *TaskManager) ResetCron() {
	tm.options.cronMu.Lock()
	defer tm.options.cronMu.Unlock()
	if tm.options.cronService != nil {
		tm.options.cronService.Stop()
	}
	tm.options.cronService = cron.New()
	tm.options.cronService.Start()
}

// Tasks will return the list of tasks
func (tm *TaskManager) Tasks() map[string]*taskq.Task {
	tm.options.taskq.tasksMu.RLock()
	defer tm.options.taskq.tasksMu.RUnlock()
	// Return a copy to prevent external modification
	tasks := make(map[string]*taskq.Task, len(tm.options.taskq.tasks))
	for k, v := range tm.options.taskq.tasks {
		tasks[k] = v
	}
	return tasks
}

// Factory will return the factory that is set
func (tm *TaskManager) Factory() Factory {
	if tm.options == nil || tm.options.taskq == nil {
		return FactoryEmpty
	}
	if tm.options.taskq.config.Redis != nil {
		return FactoryRedis
	}
	return FactoryMemory
}
