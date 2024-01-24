FROM golang:1.21-alpine

WORKDIR /2k-blacktop

RUN apk update && apk upgrade

RUN apk add --no-cache sqlite

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go run ./cmd/create_db/create_db.go

ENV PORT=8080

EXPOSE 8080

RUN go build -v .

CMD ["./2k-blacktop"]