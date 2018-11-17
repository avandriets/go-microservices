package model

import (
	"../messages"
	"context"
	"google.golang.org/grpc"
)

var connection *grpc.ClientConn
var client *messages.ContactServiceClient
var cont *context.Context

func SendContext(c *context.Context) {
	cont = c
}

func SetConnection(con *grpc.ClientConn) {
	connection = con
}

func SetClient(c *messages.ContactServiceClient) {
	client = c
}

func GetClient() messages.ContactServiceClient {
	return *client
}

func GetConnection() *grpc.ClientConn {
	return connection
}

func SendContact(contact *messages.Contact) (*messages.Contact, error) {
	//client := messages.NewContactServiceClient(connection)

	// Contact the storage-service and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()

	r, err := (*client).AddContact(*cont, &messages.ContactRequest{Contact: contact})

	return r.Contact, err
}
