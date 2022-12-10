FROM golang:alpine

RUN mkdir service

COPY . /service

WORKDIR /service

RUN go mod download

RUN go build -o main /cmd/app/main.go

CMD ./main

EXPOSE 9000