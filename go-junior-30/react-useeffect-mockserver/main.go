package main

import (
	"encoding/json"
	"net/http"
)

const port = ":8886"

func main() {

	http.HandleFunc("/mock", mockData)

	http.ListenAndServe(port, nil)
}

func mockData(w http.ResponseWriter, r *http.Request) {
	mobs := []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Job    string `json:"job"`
		Level  string `json:"level"`
		Weapon string `json:"weapon"`
	}{
		{
			Id:     1,
			Name:   "Yale",
			Job:    "Knight",
			Level:  "10",
			Weapon: "Spear",
		},
		{
			Id:     2,
			Name:   "Bubble",
			Job:    "Sword Man",
			Level:  "20",
			Weapon: "Sword",
		},
		{
			Id:     3,
			Name:   "Jack",
			Job:    "Berserker",
			Level:  "30",
			Weapon: "BareHand",
		},
		{
			Id:     4,
			Name:   "Paker",
			Job:    "Assassin",
			Level:  "40",
			Weapon: "Dagger",
		},
		{
			Id:     5,
			Name:   "Jiar",
			Job:    "Hunter",
			Level:  "50",
			Weapon: "Bow",
		},
		{
			Id:     6,
			Name:   "Wa Brother",
			Job:    "Magician",
			Level:  "60",
			Weapon: "Staff",
		},
		{
			Id:     7,
			Name:   "Sam",
			Job:    "Advanture",
			Level:  "70",
			Weapon: "Hammer",
		},
		{
			Id:     8,
			Name:   "Kasper",
			Job:    "Scientist",
			Level:  "80",
			Weapon: "Books",
		},
	}

	// set headers with allow cors and application/json
	w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your allowed origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write data to client
	err := json.NewEncoder(w).Encode(mobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
