package streaminglog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"
	"sync"
	"time"

	"github.com/equinor/radix-cli/pkg/utils/log"
)

type GetReplicasFunc[T fmt.Stringer] func() ([]T, error)

type GetLogFunc[T fmt.Stringer] func(ctx context.Context, item T, print func(text string)) error

type streamingReplicas[T fmt.Stringer] struct {
	output      io.Writer
	colorIndex  int
	getReplicas GetReplicasFunc[T]
	getLogs     GetLogFunc[T]
}

func New[T fmt.Stringer](output io.Writer, getReplicas GetReplicasFunc[T], getLogs GetLogFunc[T]) *streamingReplicas[T] {
	return &streamingReplicas[T]{
		output:      output,
		getReplicas: getReplicas,
		colorIndex:  0,
		getLogs:     getLogs,
	}
}

func (c *streamingReplicas[T]) StreamLogs(ctx context.Context) error {
	mutex := sync.Mutex{}
	syncingReplicas := make([]string, 0)
	wg := sync.WaitGroup{}

	for {
		componentReplicas, err := c.getReplicas()
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

		select {
		case <-ctx.Done():
			wg.Wait()
			return nil
		case <-time.Tick(5 * time.Second):
			continue // continue the for loop and refresh the replicas
		}
	}

}

func (c *streamingReplicas[T]) streamLogs(ctx context.Context, item T) error {
	c.colorIndex++
	color := log.GetColor(c.colorIndex)
	err := c.getLogs(ctx, item, func(text string) {
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
