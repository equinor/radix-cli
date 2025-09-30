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

type EventCallback func(Event)

const ContentTypeEventStream = "text/event-stream"

func NewEventSourceConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data any) error {
		if reader == nil {
			return errors.New("EventStreamConsumer requires a reader") // early exit
		}
		if data == nil {
			return errors.New("nil destination for EventStreamConsumer")
		}

		callback, ok := data.(EventCallback)
		if !ok {
			return errors.New("EventStreamConsumer requires a EventCallback function as destination")
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

			callback(Event{
				Type:    category,
				Message: data,
			})
		}

		// We must return some kind of error, so the Generated API client doesnt try to convert the nil-data to a json structure
		return io.EOF
	})
}

func NewEventSourceClientOptions(callback EventCallback) func(co *runtime.ClientOperation) {
	return func(co *runtime.ClientOperation) {
		// We will fallback to the old reader if the content-type is not event-stream
		oldReader := co.Reader

		co.Reader = runtime.ClientResponseReaderFunc(func(cr runtime.ClientResponse, c runtime.Consumer) (any, error) {
			ct := cr.GetHeader("Content-Type")
			mt, _, _ := mime.ParseMediaType(ct)
			if mt != ContentTypeEventStream {
				return oldReader.ReadResponse(cr, c)
			}

			err := c.Consume(cr.Body(), callback)
			return nil, err
		})
	}
}
