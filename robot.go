package main

import (
    "math/rand"
    "time"
    "fmt"
)

type Robot struct{
    ID int
    Location LatLong
}

func (r *Robot) Instruct(instruction Instruction, report chan string) {
    r.Move(instruction.Location)
    if r.IsNearTheTube() {
        report <- fmt.Sprintf("%v, %v, %v, %v", r.ID, instruction.Time.Format("15:04:05"), r.Location, traffic())
    }
}

const speed int64 = 500000 // Meters per hour (500kph), pretty quick

func (r *Robot) Move(newLocation LatLong) {
    // Robots move at constant speed in a staright line
    startingLocation := LatLong{0, 0}
    if r.Location != startingLocation { // Assume it takes no time to move to starting location
        time.Sleep(time.Duration(int64(r.Location.DistanceFrom(newLocation)) * int64(time.Hour) / speed))
    }
    r.Location = newLocation
}

func (r *Robot) IsNearTheTube() bool {
    for _, station := range tubeStations { // Seems like a dreadful way to search for a nearby station
        if r.Location.DistanceFrom(station.Location) <= 350.0 {
            return true
        }
    }
    return false
}

func traffic() string {
    switch rand.Intn(3) {
    case 0:
        return "HEAVY"
    case 1:
        return "MODERATE"
    case 2:
        return "LIGHT"
    }
    return ""
}
