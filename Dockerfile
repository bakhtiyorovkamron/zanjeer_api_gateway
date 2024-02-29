FROM golang:alpine

WORKDIR /app

COPY . .

RUN go test -bench=.

EXPOSE 7777

CMD [ "go","run","cmd/main.go" ]