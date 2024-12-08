package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetResponse() error {
	url := "http://10.128.0.2:9977/v1/user/check"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Print the raw response body
	fmt.Println(string(body))
	return nil
}

func main() {
	err := GetResponse()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
