FROM alpine:3.21.3 AS builder

RUN apk add --no-cache \
    go=1.23.8-r0 \
    make \
    gcc \
    musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server -trimpath ./cmd/app/main.go

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /app/server .

RUN addgroup appgroup && \
    adduser -S -D -H -G appgroup appuser

RUN chown -R appuser:appgroup /app
USER appuser

CMD ["./server"]