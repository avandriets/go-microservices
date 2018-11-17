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
	rpcConn, err := connectToGRPCServer()
	defer rpcConn.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/read-csv-file", controller.CSVReader).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}

func connectToGRPCServer() (*grpc.ClientConn, error) {

	serverPort, okP := os.LookupEnv("TC_SERVER_PORT")
	serverHost, okS := os.LookupEnv("TC_SERVER_HOST")

	if okS && okP {
		address = serverHost + ":" + serverPort
	}

	// Set up a connection to the storage-service.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	model.SetConnection(conn)

	return conn, err
}
