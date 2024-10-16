package main

import (
	"fmt"
	"net/http"
	clientRequest "second-task/task2/internal/client"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 100 * time.Second,
	}
	version, err := clientRequest.RequestVersion(&client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(version)

	outputString, err := clientRequest.DecodeRequest(&client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputString)

	result, statusCode, err := clientRequest.HardOpRequest(&client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result, statusCode)
}
