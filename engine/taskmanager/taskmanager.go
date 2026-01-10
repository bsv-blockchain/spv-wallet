/*
Package taskmanager is the task/job management service layer for concurrent and asynchronous tasks with cron scheduling.
*/
package taskmanager

import (
	"context"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/vmihailenco/taskq/v3"

	"github.com/bsv-blockchain/spv-wallet/engine/logging"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
)

type (

	// TaskManager implements the TaskEngine interface
	TaskManager struct {
		options *options
	}

	options struct {
		cronService *cron.Cron      // Internal cron job client
		logger      *zerolog.Logger // Internal logging
		taskq       *taskqOptions   // All configuration and options for using TaskQ
	}

	// taskqOptions holds all the configuration for the TaskQ engine
	taskqOptions struct {
		config  *taskq.QueueOptions    // Configuration for the TaskQ engine
		queue   taskq.Queue            // Queue for TaskQ
		tasks   map[string]*taskq.Task // Registered tasks
		queueMu sync.Mutex             // Mutex for queue operations to prevent race conditions
		tasksMu sync.RWMutex           // Mutex for tasks map operations
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

		// Stop the cron scheduler
		if tm.options.cronService != nil {
			tm.options.cronService.Stop()
			tm.options.cronService = nil
		}

		// Close the taskq queue
		if err := tm.options.taskq.queue.Close(); err != nil {
			return spverrors.Wrapf(err, "failed to close taskq queue")
		}

		// Empty all values and reset
		tm.options.taskq.config = nil
		tm.options.taskq.queue = nil
	}

	return nil
}

// ResetCron will reset the cron scheduler and all loaded tasks
func (tm *TaskManager) ResetCron() {
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
