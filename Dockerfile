

FROM golang:1.16-alpine

WORKDIR /app

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake 

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

EXPOSE 8080

CMD [ "go", "run", "main.go" ]