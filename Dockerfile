FROM golang:alpine

WORKDIR /app

COPY . .

EXPOSE 7777

CMD [ "go","run","cmd/main.go" ]