# Build the binary
FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./easydns-sync-ip ./cmd/main.go

# Final output image
FROM scratch
COPY --from=0 /app/easydns-sync-ip /app/easydns-sync-ip
CMD ["/app/easydns-sync-ip"]
