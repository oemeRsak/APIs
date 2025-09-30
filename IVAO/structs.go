package main

import "time"

type Pilot struct {
	ID              int          `json:"id"`
	UserID          int          `json:"userId"`
	Callsign        string       `json:"callsign"`
	ServerID        string       `json:"serverId"`
	SoftwareTypeID  string       `json:"softwareTypeId"`
	SoftwareVersion string       `json:"softwareVersion"`
	Rating          int          `json:"rating"`
	CreatedAt       string       `json:"createdAt"` // could be time.Time if properly formatted
	Time            int          `json:"time"`
	PilotSession    PilotSession `json:"pilotSession"`
	LastTrack       LastTrack    `json:"lastTrack"`
	FlightPlan      FlightPlan   `json:"flightPlan"`
}

// PilotSession contains simulator and texture information for a pilot.
type PilotSession struct {
	SimulatorID string `json:"simulatorId"`
	TextureID   int    `json:"textureId"`
}

// LastTrack holds the latest tracking data for a pilot's flight.
type LastTrack struct {
	Altitude           int       `json:"altitude"`
	AltitudeDifference int       `json:"altitudeDifference"`
	ArrivalDistance    float64   `json:"arrivalDistance"`
	DepartureDistance  float64   `json:"departureDistance"`
	GroundSpeed        int       `json:"groundSpeed"`
	Heading            int       `json:"heading"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	OnGround           bool      `json:"onGround"`
	State              string    `json:"state"`
	Timestamp          time.Time `json:"timestamp"`
	Transponder        int       `json:"transponder"`
	TransponderMode    string    `json:"transponderMode"`
	Time               int       `json:"time"`
}

// FlightPlan contains the details of a pilot's filed flight plan.
type FlightPlan struct {
	ID                       int       `json:"id"`
	Revision                 int       `json:"revision"`
	AircraftID               string    `json:"aircraftId"`
	AircraftNumber           int       `json:"aircraftNumber"`
	DepartureID              string    `json:"departureId"`
	ArrivalID                string    `json:"arrivalId"`
	AlternativeID            string    `json:"alternativeId"`
	Alternative2ID           *string   `json:"alternative2Id"` // nullable
	Route                    string    `json:"route"`
	Remarks                  string    `json:"remarks"`
	Speed                    string    `json:"speed"`
	Level                    string    `json:"level"`
	FlightRules              string    `json:"flightRules"`
	FlightType               string    `json:"flightType"`
	EET                      int       `json:"eet"`
	Endurance                int       `json:"endurance"`
	DepartureTime            int       `json:"departureTime"`
	ActualDepartureTime      *int      `json:"actualDepartureTime"` // nullable
	PeopleOnBoard            int       `json:"peopleOnBoard"`
	CreatedAt                time.Time `json:"createdAt"`
	Aircraft                 Aircraft  `json:"aircraft"`
	AircraftEquipments       string    `json:"aircraftEquipments"`
	AircraftTransponderTypes string    `json:"aircraftTransponderTypes"`
}

// Aircraft describes the aircraft used in a flight plan.
type Aircraft struct {
	IcaoCode       string `json:"icaoCode"`
	Model          string `json:"model"`
	WakeTurbulence string `json:"wakeTurbulence"`
	IsMilitary     bool   `json:"isMilitary"`
	Description    string `json:"description"`
}
