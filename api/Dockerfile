FROM golang:1.25 as builder

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .

COPY api/. .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/binary ./cmd/

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/binary /binary

ENTRYPOINT ["/binary"]