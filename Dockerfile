FROM alpine:3.16.2
RUN apk update && apk add ca-certificates && apk add bash && rm -rf /var/cache/apk/*

RUN addgroup -S -g 1000 radix && adduser -S -u 1000 -G radix radix

WORKDIR /app
COPY rx /app/rx
USER 1000
ENTRYPOINT ["/app/rx"]