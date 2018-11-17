package controller

import (
	"../messages"
	"../model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type TimeSpent struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	Elapsed string `json:"elapsed"`
}

type TestRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
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

	serverUrl, ok := os.LookupEnv("TC_GRPC_SERVER_URL")

	address := "localhost:50051"
	if ok {
		address = serverUrl
	}

	for i := 0; i < 100000; i++ {
		jsonData := TestRequest{i, "Alex", "Alex@gmail.com", "00033354"}
		jsonValue, _ := json.Marshal(jsonData)

		response, err := http.Post("http://"+address+"/post-test", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)

	timeSpent := TimeSpent{start.String(), end.String(), elapsed.String()}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&timeSpent)

	if err != nil {
		log.Printf("HTTP %s", err)
	}
}
