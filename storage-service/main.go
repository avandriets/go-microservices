package main

import (
	"context"
	"fmt"
	"go-microservices/storage-service/controller"
	"go-microservices/storage-service/messsages"
	"go-microservices/storage-service/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

import _ "github.com/lib/pq"
import "database/sql"

const (
	port = ":50051"
)

type server struct{}

func main() {
	config := model.Config{}
	config.Init()
	model.SetConfig(&config)

	db := connectToDatabase()
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	messages.RegisterContactServiceServer(s, &server{})
	// Register reflection service on gRPC storage-service.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (c *server) AddContact(ctx context.Context, in *messages.ContactRequest) (*messages.ContactResponse, error) {
	if contact, err := controller.AddContactToDB(in.Contact); err != nil {
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
