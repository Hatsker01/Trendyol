FROM golang:1.18rc1-alpine3.15
RUN mkdir post_service
COPY . /post_service
WORKDIR /post_service
# RUN apk add curl 
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
# COPY --from=builder /post_service/migrate ./migrate
# COPY migrations ./migrations
RUN go mod tidy
RUN go build -o main cmd/trend/main.go
CMD ./main
EXPOSE 8000