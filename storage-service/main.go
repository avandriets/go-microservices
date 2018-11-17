package main

import (
	"./messsages"
	"./model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
)

import _ "github.com/lib/pq"
import "database/sql"

const (
	port = ":50051"
)

func main() {
	//config := model.Config{}
	//config.Init()
	//model.SetConfig(&config)
	//
	//db := connectToDatabase()
	//defer db.Close()

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

type server struct{}

func (c *server) AddContact(ctx context.Context, in *messages.ContactRequest) (*messages.ContactResponse, error) {
	return &messages.ContactResponse{Contact: in.Contact}, nil
}

func (c *server) AddAllContacts(stream messages.ContactService_AddAllContactsServer) error {
	for {
		contact, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		log.Println(*contact.Contact)
		stream.Send(&messages.ContactResponse{Contact: contact.Contact})
	}
	return nil
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
