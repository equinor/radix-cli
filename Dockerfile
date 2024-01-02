FROM golang:1.21-alpine3.18 as builder

ENV GO111MODULE=on

RUN apk update && \
    apk add bash jq alpine-sdk sed gawk git ca-certificates curl mc && \
    apk add --no-cache gcc musl-dev

WORKDIR /app

# Install project dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy project code
COPY . /app

RUN addgroup -S -g 1000 radix && adduser -S -u 1000 -G radix radix

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o ./rootfs/rx

## Run operator
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/rootfs/rx /usr/local/bin/rx
USER 1000
ENTRYPOINT ["/usr/local/bin/rx"]

