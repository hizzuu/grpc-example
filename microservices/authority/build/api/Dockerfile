# dev
FROM golang:1.17.7-alpine as dev

WORKDIR /app

ENV GO111MODULE on
ENV CGO_ENABLED 0

RUN apk add --no-cache alpine-sdk git tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "Asia/Tokyo" > /etc/timezone

RUN go install github.com/cespare/reflex@latest

COPY . .

CMD reflex -r '(\.go$|go\.mod)' -s go run ./cmd/api/main.go

# builder
FROM golang:1.17.2-alpine as builder

WORKDIR /src

RUN apk add --no-cache alpine-sdk git tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "Asia/Tokyo" > /etc/timezone

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

# prod
FROM alpine as prod

WORKDIR /app

COPY --from=builder /src/main .

CMD ["./main"]