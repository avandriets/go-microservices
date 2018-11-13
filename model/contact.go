package model

import "fmt"

type ContactInfo struct {
	Id    int
	Name  string
	Email string
	Phone string
}

func (c *ContactInfo) convertData(row string) {
	fmt.Printf("[ row = %s ]\n", row)
}

func (c *ContactInfo) SendData(row string) {
	c.convertData(row)
}
