package main

import (
	"fmt"
	"encoding/json"
	"log/slog"
	"os"

	cfg "github.com/jvcorredor/agent-do/internal/config"
	repo "github.com/jvcorredor/agent-do/repository"
	task "github.com/jvcorredor/agent-do/task"
)

var config cfg.Config

func init() {
	if err := cfg.Initialize(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main is broad orchestration logic right now.
// Will get split out into something else, like a dedicated runner
func main() {
	logger := slog.Default()

	r := repo.New(config.Remote)
	dir, err := r.Clone()
	if err != nil {
		logger.Error("a fatal error occurred", "error", err)
		os.Exit(1)
	}

	logger.Info("repository clone successful", "dir", dir)

	if err := r.Checkout(config.Sha); err != nil {
		logger.Error("a fatal error occurred", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := r.Cleanup(); err != nil {
			logger.Error("cleanup failed", "error", err)
			os.Exit(1)
		}
	}()

	t := task.New(task.WithLogger(logger))
	if err := run(logger, t); err != nil {
		logger.Error("failed", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(logger *slog.Logger, t *task.Task) error {
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	logger.Info("run concluded", "task", string(bytes))
	return nil
}
