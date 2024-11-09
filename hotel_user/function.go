package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CheckUserResponse struct {
	Correct bool `json:"correct"`
}

func CheckUser(username, password string) error {
	baseURL := "http://10.128.0.2:8087/v1/user/check"
	params := url.Values{}
	params.Add("username", username)
	params.Add("password", password)
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())


	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	defer resp.Body.Close()

	var result CheckUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error prase request: %v", err)
	}

	if result.Correct {
		fmt.Println("successfully login")
	} else {
		fmt.Println("failed to login")
	}

	return nil
}

func main() {
	err := CheckUser("Cornell_30", "30303030303030303030")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
