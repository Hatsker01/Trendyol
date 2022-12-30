FROM golang:1.18rc1-alpine3.15
RUN mkdir post_service
COPY . /post_service
WORKDIR /post_service
RUN go mod download
RUN go build -o main cmd/trend/main.go
CMD ./main
EXPOSE 8000