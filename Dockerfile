FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=LINUX GOARCH=amd64 go build -a -installsuffix cgo -o /app/bin/main .

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
	tzdata

ENV TZ UTC

COPY --from=bilder /app/bin/main .

EXPOSE 8090

CMD ["./main", "rest", "-p", "8070"]
