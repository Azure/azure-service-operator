// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"context"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/go-logr/logr"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"
)

type Watcher struct {
	path string

	ctx    context.Context
	cancel context.CancelFunc
	logger logr.Logger
}

func NewWatcher(logger logr.Logger) *Watcher {
	logger = logger.WithName("configwatcher")

	return &Watcher{
		path:   MountPath,
		logger: logger,
	}
}

func (w *Watcher) Start(ctx context.Context) error {
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.logger.V(Info).Info("starting watcher", "path", w.path)
	go w.runImpl()

	return nil
}

func (w *Watcher) Stop() {
	w.logger.V(Info).Info("stopping watcher", "path", w.path)
	w.cancel()
}

func (w *Watcher) runImpl() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		w.logger.Error(err, "failed to create fsnotify watcher")
		return
	}
	defer watcher.Close()

	err = watcher.Add(w.path)

	for {
		var ok bool
		var event fsnotify.Event

		select {
		case <-w.ctx.Done():
			w.logger.Error(err, "Aborting watcher")
			return
		case event, ok = <-watcher.Events:
			if !ok {
				w.logger.Error(err, "Watcher events channel closed")
				return
			}
			if event.Has(fsnotify.Write | fsnotify.Create | fsnotify.Remove | fsnotify.Rename) {
				w.logger.V(Info).Info("Secret updated, exiting...")
				// TODO: Ideally we would give up the lease here too
				os.Exit(0)
			}
		case err, ok = <-watcher.Errors:
			if !ok {
				w.logger.Error(err, "Watcher error channel closed")
				return
			}
		}
	}
}
