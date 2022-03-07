package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	path := "./msg_array.json"
	msgList := Parse(path)
	fmt.Printf("msg_list len: %d \n", len(msgList))
}

func TestStreamParse(t *testing.T) {
	path := "./msg_array.json"
	msgList := StreamParse(path)
	fmt.Printf("msg_list len: %d \n", len(msgList))
}
