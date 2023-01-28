# Trendyol Backend #

## Running 
    docker-compose build

    docker-compose up
## Check
    http://hostname:8090/swagger/index.html#

## Usage for API
    github.com/gin-gonic/gin  package


All rows ended with **swagger** and there documintation for API
 
This project example for **https://www.trendyol.com/** site

**Used postgresql database**

There are three microservices. They call each other with gRPC ***google.golang.org/grpc***.

All services run easIly with docker-compose file
