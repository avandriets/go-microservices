package main

import (
	"./controller"
	"./messsages"
	"./model"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

import _ "github.com/lib/pq"
import "database/sql"

const (
	port = ":50051"
)

func main() {
	r := mux.NewRouter()
	// Just for testing purpose
	r.HandleFunc("/post-test", controller.JsonRequest4Test).Methods("POST")

	if err := http.ListenAndServe(":50051", r); err != nil {
		log.Fatal(err)
	}
	//config := model.Config{}
	//config.Init()
	//model.SetConfig(&config)
	//
	//db := connectToDatabase()
	//defer db.Close()
	//
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//messages.RegisterContactServiceServer(s, &server{})
	//// Register reflection service on gRPC storage-service.
	//reflection.Register(s)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}

type server struct{}

func (c *server) AddContact(ctx context.Context, in *messages.ContactRequest) (*messages.ContactResponse, error) {
	if contact, err := controller.AddContact(in.Contact); err != nil {
		return &messages.ContactResponse{Contact: contact}, err
	}
	return &messages.ContactResponse{Contact: &messages.Contact{Id: -1}}, nil
}

func connectToDatabase() *sql.DB {
	config := model.GetConfig()
	connStr := config.GetDBConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable connect to DB: %v", err))
	}

	model.SetDatabase(db)

	return db
}
