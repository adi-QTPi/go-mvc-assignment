
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# ---- Run Stage ----
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/pkg/views ./pkg/views 
COPY --from=builder /app/public ./public
COPY --from=builder /app/uploads ./uploads

# Run the binary
CMD ["./main"]
