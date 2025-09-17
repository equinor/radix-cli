package streaminglog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/component"
	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/go-openapi/strfmt"
)

type GetReplicasFunc func() (map[string][]string, error)

type streamingReplicas struct {
	output      io.Writer
	apiClient   *radixapi.Radixapi
	appName     string
	since       strfmt.DateTime
	previous    string
	colorIndex  int
	getReplicas GetReplicasFunc
}

func NewStreamingReplicas(apiClient *radixapi.Radixapi, output io.Writer, appName string, since time.Duration, previousLog bool, getReplicas GetReplicasFunc) *streamingReplicas {
	now := time.Now()
	sinceTime := now.Add(-since)
	sinceDt := strfmt.DateTime(sinceTime)
	previous := strconv.FormatBool(previousLog)

	return &streamingReplicas{
		apiClient:   apiClient,
		output:      output,
		appName:     appName,
		since:       sinceDt,
		previous:    previous,
		getReplicas: getReplicas,
		colorIndex:  0,
	}
}

func (c *streamingReplicas) StreamLogs(ctx context.Context) error {
	mutex := sync.Mutex{}
	syncingReplicas := make([]string, 0)
	wg := sync.WaitGroup{}

	for {
		componentReplicas, err := c.getReplicas()
		if err != nil {
			return err
		}

		for componentName, replicas := range componentReplicas {
			for _, replica := range replicas {
				if slices.Contains(syncingReplicas, replica) {
					continue
				}

				wg.Go(func() {
					mutex.Lock()
					syncingReplicas = append(syncingReplicas, replica)
					mutex.Unlock()
					if err := c.startReplicaStream(ctx, componentName, replica); err != nil {
						mutex.Lock()
						syncingReplicas = slices.DeleteFunc(syncingReplicas, func(item string) bool { return item == replica })
						mutex.Unlock()
					}
				})

			}
		}

		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		case <-time.Tick(15 * time.Second):
			continue // continue the for loop and refresh the replicas
		}
	}

}

func (c *streamingReplicas) startReplicaStream(ctx context.Context, componentName, replica string) error {
	c.colorIndex++
	color := log.GetColor(c.colorIndex)

	logParameters := component.NewLogParamsWithContext(ctx)
	logParameters.WithAppName(c.appName)
	logParameters.WithDeploymentName("irrelevant")
	logParameters.WithComponentName(componentName)
	logParameters.WithPodName(replica)
	logParameters.WithFollow(pointers.Ptr("true"))
	logParameters.SetSinceTime(&c.since)
	logParameters.WithPrevious(&c.previous)

	resp, err := c.apiClient.Component.Log(logParameters, nil, consumer.NewEventSourceClientOptions(func(event consumer.Event) {
		switch event.Type {
		case "event":
			switch event.Message {
			case "started":
				log.PrintLine(c.output, replica, "stream started...", color)
			case "completed":
				log.PrintLine(c.output, replica, "stream closed.", color)
			}
		case "data":
			log.PrintLine(c.output, replica, event.Message, color)
		}
	}))
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		log.PrintLine(c.output, replica, log.Red(fmt.Sprintf("error: %s", err.Error())), color)
		return err
	}

	lines := strings.Split(resp.Payload, "\n")
	for _, line := range lines {
		log.PrintLine(c.output, replica, line, color)
	}
	log.PrintLine(c.output, replica, "stream closed.", color)

	return nil
}
