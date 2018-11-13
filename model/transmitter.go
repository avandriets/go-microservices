package model

type Transmitter interface {
	convertData(row string)
	SendData(row string)
}
