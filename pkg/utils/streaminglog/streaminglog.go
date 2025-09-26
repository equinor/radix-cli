package streaminglog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"
	"sync"
	"time"

	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/go-openapi/runtime"
)

// GetReplicasFunc is a function type that returns a list of items (replicas) to stream logs from, a boolean indicating if we are finished, and an error if any occurred.
type GetReplicasFunc[T fmt.Stringer] func() ([]T, bool, error)

type GetLogFunc[T fmt.Stringer] func(ctx context.Context, item T, since time.Time, print func(text string)) error

type streamingReplicas[T fmt.Stringer] struct {
	output      io.Writer
	colorIndex  int
	since       time.Time
	getReplicas GetReplicasFunc[T]
	getLogs     GetLogFunc[T]
}

var loopDuration = 2 * time.Second

func New[T fmt.Stringer](output io.Writer, getReplicas GetReplicasFunc[T], getLogs GetLogFunc[T], since time.Duration) *streamingReplicas[T] {
	sinceTime := time.Now().Add(-since)

	return &streamingReplicas[T]{
		output:      output,
		getReplicas: getReplicas,
		colorIndex:  0,
		since:       sinceTime,
		getLogs:     getLogs,
	}
}

func (c *streamingReplicas[T]) StreamLogs(ctx context.Context) error {
	mutex := sync.Mutex{}
	syncingReplicas := make([]string, 0)
	wg := sync.WaitGroup{}

	for {
		componentReplicas, finished, err := c.getReplicas()
		if err != nil {
			return err
		}

		for _, item := range componentReplicas {
			if slices.Contains(syncingReplicas, item.String()) {
				continue
			}

			wg.Go(func() {
				mutex.Lock()
				syncingReplicas = append(syncingReplicas, item.String())
				mutex.Unlock()
				if err := c.streamLogs(ctx, item); err != nil && !errors.Is(err, io.EOF) {
					mutex.Lock()
					syncingReplicas = slices.DeleteFunc(syncingReplicas, func(i string) bool { return i == item.String() })
					mutex.Unlock()
				}
			})
		}

		// If we are finished, dont loop again
		if finished {
			wg.Wait()
			return nil
		}

		// Wait for either context cancellation or loop duration to elapse
		c.since = time.Now().Add(-loopDuration)
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
	color := log.GetColor(c.colorIndex)
	err := c.getLogs(ctx, item, c.since, func(text string) {
		log.PrintLine(c.output, item.String(), text, color)
	})
	if err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
			return nil
		}

		log.PrintLine(c.output, item.String(), log.Red(fmt.Sprintf("error: %s", err.Error())), color)
		return err
	}
	return nil
}

func CreateLogStreamer(print func(text string)) func(co *runtime.ClientOperation) {
	return consumer.NewEventSourceClientOptions(func(event consumer.Event) {
		switch event.Type {
		case "event":
			switch event.Message {
			case "started":
				print("stream started...")
			case "completed":
				print("stream closed.")
			}
		case "data":
			print(event.Message)
		}
	})
}
