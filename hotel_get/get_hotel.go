package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hotel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func GetHotels() ([]Hotel, error) {
	baseURL := "http://localhost:9977/v1/hotels"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed request: %v", err)
	}
	defer resp.Body.Close()

	var hotels []Hotel
	if err := json.NewDecoder(resp.Body).Decode(&hotels); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return hotels, nil
}

func main() {
	hotels, err := GetHotels()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, hotel := range hotels {
		fmt.Printf("Hotel: %s, Location: %s\n", hotel.Name, hotel.Location)
	}
}
