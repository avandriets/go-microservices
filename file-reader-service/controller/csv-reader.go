package controller

import (
	"fmt"
	"go-microservices/file-reader-service/blob-reader"
	"go-microservices/file-reader-service/messages"
	"go-microservices/file-reader-service/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

func CSVReader(w http.ResponseWriter, r *http.Request) {
	blob_reader.ReadCSVFile("/Users/andriets/go/src/go-microservices/file-reader-service/data/data.csv", dataSender)
}
