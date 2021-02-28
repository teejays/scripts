FROM golang:1.13.4-alpine3.10

LABEL maintainer="hello@teejay.me"
ENV GO111MODULE=on

RUN apk update && apk add make git

RUN mkdir -p app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o power-n main.go

CMD ["./power-n"]