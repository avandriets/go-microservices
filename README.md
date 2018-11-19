# Go Microservices GRPC

This repo illustrate how to build application that based on microservices architecture with go and gRpc
There are two services
1. file-reader-service (read CSV files and send data to other server to save data)
2. storage-service (get requests from other services by GRPC protocol insert data into DB and send response back)

All services and infrastructure stuff run in docker

## Run project (please use docker)

```
$ docker-compose up --build or just docker-compose up
```

## Endpoints
* http://localhost:9000/read-csv-file (GET) initiate processing csv file that stored inside container

I understand that I should use POST method and upload file to the server and then processing it but it is just for
test purpose to initiate processing

* if you want start services not in docker you should do something with data file go-microservices/data/data.csv