FROM golang:1.19.6-alpine as build-env
WORKDIR /micro

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o micro ./cmd

EXPOSE 8283

CMD ["./micro"]