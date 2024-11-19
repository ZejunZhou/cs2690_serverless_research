package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReservationRequest struct {
	HotelID  string `json:"hotel_id"`
	UserID   string `json:"user_id"`
	Duration int    `json:"duration"`
}

type ReservationResponse struct {
	Status string `json:"status"`
}

func PostReservation(req ReservationRequest) (*ReservationResponse, error) {
	baseURL := "http://localhost:9977/v1/reservations"
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding request: %v", err)
	}

	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed request: %v", err)
	}
	defer resp.Body.Close()

	var response ReservationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &response, nil
}

func main() {
	req := ReservationRequest{
		HotelID:  "123",
		UserID:   "456",
		Duration: 3,
	}

	res, err := PostReservation(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Reservation Status: %s\n", res.Status)
}
