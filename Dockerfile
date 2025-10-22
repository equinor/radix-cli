FROM --platform=$BUILDPLATFORM docker.io/golang:1.25-alpine3.22 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

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
RUN go build -ldflags "-s -w" -a -installsuffix cgo -o ./rootfs/rx ./cli/rx

## Run operator
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/rootfs/rx /usr/local/bin/rx
USER 1000
ENTRYPOINT ["/usr/local/bin/rx"]

