package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Clients holds a list of pilots from the IVAO API response.
type Clients struct {
	Pilots []Pilot `json:"pilots"`
}

// WhazzupResponse is the top-level structure for the IVAO tracker API response.
type WhazzupResponse struct {
	Clients Clients `json:"clients"`
}

// GetFlightsFromIVAO fetches flights from the IVAO API and filters them by the given airport.
// It appends new flights to the processed slice if not already present.
func GetFlightsFromIVAO(airport string, processed *[]Pilot) {
	resp, err := http.Get("https://api.ivao.aero/v2/tracker/whazzup")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("HTTP Status: " + resp.Status)
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var whazzup WhazzupResponse

	if err := json.Unmarshal(byt, &whazzup); err != nil {
		panic(err)
	}

	// Filter flights by airport
	for _, flight := range whazzup.Clients.Pilots {
		// Check if flight is already in processed
		alreadyProcessed := false
		for _, p := range *processed {
			if p.ID == flight.ID {
				alreadyProcessed = true
				break
			}
		}
		if !alreadyProcessed && (flight.FlightPlan.DepartureID == airport || flight.FlightPlan.ArrivalID == airport) {
			*processed = append(*processed, flight)
		}
	}

	// Print the filtered flights for debugging
	/*
		for _, f := range *processed {
			jsonBytes, _ := json.MarshalIndent(f, "", "  ")
			println(string(jsonBytes))
		}
	*/

}
