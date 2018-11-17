package controller

import (
	"../messages"
	"../model"
	"../reader"
	"encoding/json"
	"fmt"
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

func CsvReader(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	reader.ReadCsvFile("./data/data.csv", dataSender)
	end := time.Now()
	elapsed := end.Sub(start)

	timeSpent := TimeSpent{start.String(), end.String(), elapsed.String()}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&timeSpent)

	if err != nil {
		log.Printf("HTTP %s", err)
	}
}
