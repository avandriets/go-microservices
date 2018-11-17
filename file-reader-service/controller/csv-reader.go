package controller

import (
	"../messages"
	"../model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TimeSpent struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	Elapsed string `json:"elapsed"`
}

func dataSender(row string) {
	if len(row) > 0 {
		fields := strings.Split(row, ",")

		if id, err := strconv.Atoi(fields[0]); err == nil && len(fields) > 0 {
			contact := messages.Contact{Id: int32(id), Name: fields[1], Email: fields[2], Phone: fields[3]}
			fmt.Printf("[ rowSend = %s ]\n", row)
			if response, errorResp := model.SendContact(&contact); errorResp == nil {
				fmt.Printf("[ rowRecieved = %s ]\n", response)
			} else {
				log.Printf("error: %s", errorResp.Error())
			}
		}
	}
}

func sendTestRequest() {

	stream, err := model.GetClient().AddAllContacts(context.Background())
	if err != nil {
		log.Fatal("Error")
	}

	doneCh := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				doneCh <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.Contact)
		}
	}()

	for i := 0; i < 100000; i++ {
		contact := messages.Contact{Id: int32(i), Name: "Alex", Email: "alex@gmail.com", Phone: "050479888"}
		err := stream.Send(&messages.ContactRequest{Contact: &contact})

		if err != nil {
			log.Fatal("Error")
		}

		//if response, errorResp := model.SendContact(&contact); errorResp == nil {
		//	log.Printf("[ rowRecieved = %s ]\n", response)
		//} else {
		//	log.Printf("error: %s", errorResp.Error())
		//}
	}
}

func CsvReader(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//reader.ReadCsvFile("./data/data.csv", dataSender)
	sendTestRequest()
	end := time.Now()
	elapsed := end.Sub(start)

	timeSpent := TimeSpent{start.String(), end.String(), elapsed.String()}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&timeSpent)

	if err != nil {
		log.Printf("HTTP %s", err)
	}
}
