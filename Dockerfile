FROM golang:1.18.5-alpine3.16 as builder

ENV GO111MODULE=on

RUN apk update && \
    apk add bash jq alpine-sdk sed gawk git ca-certificates curl mc && \
    apk add --no-cache gcc musl-dev
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.3.3 && \
    go install github.com/rakyll/statik@v0.1.7 && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.2

WORKDIR /app

# Install project dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy project code
COPY . /app

# lint and unit tests
RUN staticcheck ./... && \
    go vet ./... && \
    CGO_ENABLED=0 GOOS=linux go test ./...

# Build
WORKDIR /app
CMD sh
#RUN make release
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o ./rootfs/radix-cli
#RUN addgroup -S -g 1000 radix
#RUN adduser -S -u 1000 -G radix radix
#
## Run operator
#FROM scratch
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /app/rootfs/radix-cli /usr/local/bin/rx
#USER 1000
#CMD sh
##ENTRYPOINT ["/usr/local/bin/radix-operator"]
