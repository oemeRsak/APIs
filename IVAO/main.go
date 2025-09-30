package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

/*
TODO: Move Info Display to Left
TODO: Add list of Flights under Control
TODO: How to move Flights between Lists ?
*/

// ? Flights holds the list of currently tracked pilot flights.
var Flights []Pilot
var Departures []Pilot
var Arrivals []Pilot

// ? AIRPORT ICAO
var ICAO string = "EDDM"

// * Main initializes the TUI application, sets up views, keybindings, and starts the main event loop.
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	//* Main Code
	g.Mouse = true
	g.Cursor = true

	g.SetManagerFunc(layout)

	//* Periodically fetch flights from IVAO and update the Lists every second
	go func() {
		for {
			time.Sleep(time.Second)
			g.Update(func(g *gocui.Gui) error {
				GetFlightsFromIVAO(ICAO, &Flights)

				vD, err := g.View("DEL")
				if err != nil {
					return err
				}
				vA, err := g.View("APP")
				if err != nil {
					return err
				}

				vD.Clear()
				vA.Clear()

				Departures = nil
				Arrivals = nil

				//? Print each flight on a separate line of Its View
				for _, flight := range Flights {
					if flight.FlightPlan.DepartureID == ICAO {
						Departures = append(Departures, flight)
						fmt.Fprintf(vD, "Flight ID: %d, Callsign: %v\n", flight.ID, flight.Callsign)
					} else if flight.FlightPlan.ArrivalID == ICAO {
						Arrivals = append(Arrivals, flight)
						fmt.Fprintf(vA, "Flight ID: %d, Callsign: %v\n", flight.ID, flight.Callsign)
					}
				}

				//? Update TIME
				if vT, err := g.View("TIME"); err == nil {
					vT.Clear()
					utc := time.Now().UTC().Format("15:04:05")
					local := time.Now().Format("15:04:05")
					fmt.Fprintf(vT, "Zulu: %s | Local: %s", utc, local)
				}

				return nil
			})
		}
	}()

	//* MouseClick handler for departures
	g.SetKeybinding("DEL", gocui.MouseLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, cy := v.Cursor()
		if cy < len(Departures) {
			flight := Departures[cy]
			showFlightInfo(g, flight)
		}
		return nil
	})

	//* Mouse handler for departures REMOVE
	//TODO: Work in progress
	g.SetKeybinding("DEL", gocui.MouseRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, cy := v.Cursor()
		if cy < len(Departures) {
			flight := Departures[cy]
			fmt.Fprintln(v, flight.ID)
		}
		return nil
	})

	//* Mouse handler for arrivals
	g.SetKeybinding("APP", gocui.MouseLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, cy := v.Cursor()
		if cy < len(Arrivals) {
			flight := Arrivals[cy]
			showFlightInfo(g, flight)
		}
		return nil
	})

	//* Mouse handler for arrivals REMOVE
	//TODO: Work in progress
	g.SetKeybinding("APP", gocui.MouseRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, cy := v.Cursor()
		if cy < len(Arrivals) {
			flight := Arrivals[cy]
			fmt.Fprintln(v, flight.ID)
		}
		return nil
	})

	//

	//* Scroll down with mouse
	g.SetKeybinding("", gocui.MouseWheelDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ox, oy := v.Origin()
		return v.SetOrigin(ox, oy+1)
	})

	//* Scroll up with mouse
	g.SetKeybinding("", gocui.MouseWheelUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ox, oy := v.Origin()
		if oy > 0 {
			return v.SetOrigin(ox, oy-1)
		}
		return nil
	})

	//* Handle Keyboard: Exit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// * layout arranges the main views (left, mid, right) in the TUI.
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("DEL", 1, 0, maxX/4-1, maxY/3*2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Loading...")

		v.Title = "Delivery"
		v.Autoscroll = false
		v.Wrap = true
	}

	if v, err := g.SetView("APP", 1, maxY/3*2+1, maxX/4-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Loading...")

		v.Title = "Approach"
		v.Autoscroll = false
		v.Wrap = true
	}

	if v, err := g.SetView("GND", maxX/4+1, 0, maxX/4*2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintln(v, "Loading...")

		v.Title = "Ground"
		v.Autoscroll = false
		v.Wrap = false
	}

	if v, err := g.SetView("INFO", maxX/4*2+1, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Informations"
	}

	if v, err := g.SetView("CMDS", maxX/4*2+1, maxY-3, maxX/4*3-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, " X  R  Q ")

		v.Title = "Commands"
	}

	if v, err := g.SetView("TIME", maxX/4*3+1, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Time"
		v.Autoscroll = false
		v.Wrap = false
	}

	return nil
}

// Helper to display info
func showFlightInfo(g *gocui.Gui, flight Pilot) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("INFO")
		if err != nil {
			return err
		}
		v.Clear()
		jsonBytes, _ := json.MarshalIndent(flight, "", "  ")
		fmt.Fprint(v, string(jsonBytes))
		return nil
	})
}

// * quit is the keybinding handler to exit the application.
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
