FROM golang:alpine

RUN apk add gcc g++ ca-certificates --no-cache

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./core

EXPOSE 8080

ENTRYPOINT ["/app/core"]
