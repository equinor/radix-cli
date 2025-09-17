package consumer

import (
	"bufio"
	"errors"
	"io"
	"mime"
	"strings"

	"github.com/go-openapi/runtime"
)

type Event struct {
	Type    string
	Message string
}

const ContentTypeEventStream = "text/event-stream"

func NewEventSourceConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data any) error {
		if reader == nil {
			return errors.New("EventStreamConsumer requires a reader") // early exit
		}
		if data == nil {
			return errors.New("nil destination for EventStreamConsumer")
		}

		es, ok := data.(chan Event)
		if !ok {
			return errors.New("EventStreamConsumer requires a pointer to a channel of Event")
		}

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}

			pos := strings.Index(line, ":")
			if pos == -1 {
				continue
			}
			if pos == 0 {
				continue // ignore comments
			}

			category := line[0:pos]
			data := line[pos+2:]

			es <- Event{
				Type:    category,
				Message: data,
			}
		}

		if scanner.Err() != nil {
			return scanner.Err()
		}

		return io.EOF
	})
}

func NewEventSourceClientOptions(f func(event Event)) func(co *runtime.ClientOperation) {
	return func(co *runtime.ClientOperation) {
		// We will fallback to the old reader if the content-type is not event-stream
		oldReader := co.Reader

		co.Reader = runtime.ClientResponseReaderFunc(func(cr runtime.ClientResponse, c runtime.Consumer) (interface{}, error) {
			ct := cr.GetHeader("Content-Type")
			mt, _, _ := mime.ParseMediaType(ct)
			if mt != ContentTypeEventStream {
				return oldReader.ReadResponse(cr, c)
			}

			eventStream := make(chan Event, 1000)
			defer close(eventStream)
			go func() {

				for {
					select {
					case event := <-eventStream:
						f(event)
					case <-co.Context.Done():
						return
					}
				}
			}()

			err := c.Consume(cr.Body(), eventStream)
			return nil, err
		})
	}
}
