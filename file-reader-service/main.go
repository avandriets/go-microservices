package main

import (
	"./controller"
	"./model"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	address = "localhost:50051"
)

func main() {
	rpcConn, err := connectToGrpcServer()
	defer rpcConn.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	// Just for testing purpose
	r.HandleFunc("/read-csv-file", controller.CsvReader).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}

func connectToGrpcServer() (*grpc.ClientConn, error) {

	serverUrl, ok := os.LookupEnv("TC_GRPC_SERVER_URL")

	if ok {
		address = serverUrl
	}

	// Set up a connection to the storage-service.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	model.SetConnection(conn)

	return conn, err
}
