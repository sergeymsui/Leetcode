package main

import (
	"errors"
	"fmt"
)

type MessageError struct {
	name string
}

type Mesage struct {
	text string
	MessageError
}

func NewMesage() *Mesage {
	return &Mesage{
		text: "",
		MessageError: MessageError{
			name: "Empty",
		},
	}
}

var LocalError = errors.New("Local error")

func main() {
	m := NewMesage()

	fmt.Println(m.MessageError)
}
