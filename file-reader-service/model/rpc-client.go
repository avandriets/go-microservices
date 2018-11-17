package model

import (
	"../messages"
	"context"
	"google.golang.org/grpc"
	"time"
)

var (
	connection *grpc.ClientConn
	client     messages.ContactServiceClient
)

func SetClient(c *messages.ContactServiceClient) {
	client = *c
}

func SetConnection(c *grpc.ClientConn) {
	connection = c
}

func GetConnection() *grpc.ClientConn {
	return connection
}

func SendContact(contact *messages.Contact) (*messages.Contact, error) {
	// Contact the storage-service and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.AddContact(ctx, &messages.ContactRequest{Contact: contact})

	return r.Contact, err
}
