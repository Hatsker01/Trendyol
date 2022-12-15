FROM golang:alpine

RUN mkdir service

COPY . /service

WORKDIR /service

RUN go mod tidy

RUN go build -o main cmd/trend/main.go

CMD ./main

EXPOSE 9000