FROM alpine:3.15
RUN apk update && apk add ca-certificates && apk add bash && rm -rf /var/cache/apk/*

RUN addgroup -S radix-cli && adduser -S radix-cli -G radix-cli

WORKDIR /app
COPY rx /app/rx
USER radix-cli
ENTRYPOINT ["/app/rx"]