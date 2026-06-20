package task

import (
	"log/slog"
	"time"
)

type TaskStatus string

const (
	Solved     TaskStatus = "SOLVED"
	Unsolved   TaskStatus = "UNSOLVED"
	BuildError TaskStatus = "BUILD_ERROR"
	Flaky      TaskStatus = "FLAKY"
)

type Task struct {
	logger   *slog.Logger
	Status   TaskStatus    `json:"status"`
	Passed   bool          `json:"passed"`
	Duration time.Duration `json:"duration"`
}

type Option func(*Task)

func WithLogger(logger *slog.Logger) Option {
	return func(t *Task) {
		t.logger = logger
	}
}

func New(opts ...Option) *Task {
	t := Task{
		Status:   Unsolved,
		Passed:   false,
		Duration: 0,
	}

	for _, o := range opts {
		o(&t)
	}

	return &t
}
