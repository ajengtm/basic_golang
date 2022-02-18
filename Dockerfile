FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -ldflags="-s -w" -o bin/basic_golang main.go; 

EXPOSE 8089 8090

CMD [ "/basic_golang" ]