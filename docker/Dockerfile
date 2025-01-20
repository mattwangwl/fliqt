FROM golang:1.23.4-alpine3.19 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd/main.go .
COPY internal internal


RUN go build -o ./build/server


FROM alpine:3.19

# Install time zone package
RUN apk add --update-cache tzdata

# Copy binary
workdir /app
COPY --from=builder /app/build/server /app/server
COPY ./config /app/config

EXPOSE 8801

# Run binary
ENTRYPOINT ["/app/server"]
