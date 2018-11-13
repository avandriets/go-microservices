package blob_reader

import (
	"../model"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadCSVFile(fileName string, transmitter model.Transmitter) {
	const BufferSize = 100

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)
	headOfString := ""

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		bufferedString := strings.Trim(string(buffer[:bytesread]), " ")
		bufferedString = headOfString + bufferedString

		rowsArray := strings.Split(bufferedString, "\n")
		fullRow := strings.LastIndex(bufferedString, "\n") == len(bufferedString)-1

		if len(rowsArray) > 1 {
			for index, value := range rowsArray {
				if index < len(rowsArray)-1 {
					transmitter.SendData(value)
				} else if index == len(rowsArray)-1 && fullRow {
					transmitter.SendData(value)
				} else if index == len(rowsArray)-1 && !fullRow {
					headOfString = value
				}
			}
		} else if len(rowsArray) == 1 && fullRow {
			transmitter.SendData(bufferedString)
		} else if len(rowsArray) == 1 && !fullRow {
			headOfString = bufferedString
		}
	}
}
