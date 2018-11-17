package model

import (
	"context"
	"go-microservices/file-reader-service/messages"
	"google.golang.org/grpc"
	"time"
)

var connection *grpc.ClientConn

func SetConnection(database *grpc.ClientConn) {
	connection = database
}

func GetConnection() *grpc.ClientConn {
	return connection
}

func SendContact(contact *messages.Contact) (*messages.Contact, error) {
	c := messages.NewContactServiceClient(connection)

	// Contact the storage-service and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddContact(ctx, &messages.ContactRequest{Contact: contact})

	return r.Contact, err
}
