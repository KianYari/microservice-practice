FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

RUN go env -w GOPROXY=https://goproxy.io,direct

COPY gateway/go.mod gateway/go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./gateway/cmd/

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .


CMD ["./main"]