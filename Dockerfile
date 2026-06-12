FROM golang:1.26.4-alpine AS builder

WORKDIR /server

# Install tern - note the explicit path
RUN go install github.com/jackc/tern@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

# Copy tern from correct location
COPY --from=builder /go/bin/tern /usr/local/bin/tern

# Copy your app and migrations
COPY --from=builder /server/main .
COPY --from=builder /server/migrations ./migrations

# Verify tern was copied
RUN ls -la /usr/local/bin/tern

EXPOSE 3000

CMD ["./main"]