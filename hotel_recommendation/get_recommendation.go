package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Recommendation struct {
	HotelID   string `json:"hotel_id"`
	Category  string `json:"category"`
	Rating    float64 `json:"rating"`
}

func GetRecommendations() ([]Recommendation, error) {
	baseURL := "http://localhost:9977/v1/recommendations"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed request: %v", err)
	}
	defer resp.Body.Close()

	var recommendations []Recommendation
	if err := json.NewDecoder(resp.Body).Decode(&recommendations); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return recommendations, nil
}

func main() {
	recommendations, err := GetRecommendations()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, rec := range recommendations {
		fmt.Printf("Recommendation: %s (Rating: %.2f)\n", rec.Category, rec.Rating)
	}
}
