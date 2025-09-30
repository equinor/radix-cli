package replicalog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/sirupsen/logrus"
)

type Item interface {
	Identifier() string
	Created() time.Time
}

// GetReplicasFunc is a function type that returns a list of items (replicas) to stream logs from, a boolean indicating if we are finished, and an error if any occurred.
type GetReplicasFunc[T Item] func() ([]T, bool, error)

type GetLogFunc[T Item] func(ctx context.Context, item T, since time.Time, callback consumer.EventCallback) error

type streamingReplicas[T Item] struct {
	output      io.Writer
	colorIndex  int
	since       time.Time
	getReplicas GetReplicasFunc[T]
	getLogs     GetLogFunc[T]
}

var loopDuration = 2 * time.Second

func New[T Item](output io.Writer, getReplicas GetReplicasFunc[T], getLogs GetLogFunc[T], since time.Duration) *streamingReplicas[T] {
	sinceTime := time.Now().Add(-since)

	return &streamingReplicas[T]{
		output:      output,
		getReplicas: getReplicas,
		colorIndex:  0,
		since:       sinceTime,
		getLogs:     getLogs,
	}
}

func (c *streamingReplicas[T]) StreamLogs(ctx context.Context, exitOnFailure bool) error {
	mutex := sync.Mutex{}
	syncingReplicas := make([]string, 0)
	wg := sync.WaitGroup{}

	for {
		componentReplicas, finished, err := c.getReplicas()
		if err != nil && !errors.Is(err, ErrJobFailed) {
			return err
		}

		for _, item := range componentReplicas {
			if slices.Contains(syncingReplicas, item.Identifier()) {
				continue
			}

			wg.Go(func() {
				mutex.Lock()
				syncingReplicas = append(syncingReplicas, item.Identifier())
				mutex.Unlock()
				if err := c.streamLogs(ctx, item); err != nil && !errors.Is(err, io.EOF) {
					mutex.Lock()
					syncingReplicas = slices.DeleteFunc(syncingReplicas, func(i string) bool { return i == item.Identifier() })
					mutex.Unlock()
				}
			})
		}

		// If we are finished, dont loop again
		if finished {
			wg.Wait()
			if errors.Is(err, ErrJobFailed) && exitOnFailure {
				logrus.Error(err.Error())
				os.Exit(2)
			}
			return nil
		}

		select {
		case <-ctx.Done():
			wg.Wait()
			return nil
		case <-time.Tick(loopDuration):
			continue // continue the for loop and refresh the replicas
		}
	}
}

func (c *streamingReplicas[T]) streamLogs(ctx context.Context, item T) error {
	c.colorIndex++
	color := GetColor(c.colorIndex)

	since := c.since
	if item.Created().After(since) {
		since = item.Created()
	}

	err := c.getLogs(ctx, item, since, func(event consumer.Event) {
		switch event.Type {
		case "event":
			switch event.Message {
			case "started":
				PrintLine(c.output, item.Identifier(), "stream started...", color)
			case "completed":
				PrintLine(c.output, item.Identifier(), "stream closed.", color)
			}
		case "data":
			PrintLine(c.output, item.Identifier(), event.Message, color)
		}
	})
	if err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
			return nil
		}

		PrintLine(c.output, item.Identifier(), Red(fmt.Sprintf("error: %s", err.Error())), color)
		return err
	}
	return nil
}
