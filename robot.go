package main

import (
    "time"
    "fmt"
)

type Robot struct{
    ID int
    Location LatLong
}

func (r *Robot) ReceiveInstruction(instruction Instruction) {
    r.Move(instruction.Location)
}

const speed int64 = 100000 // Meters per hour (100kph)

func (r *Robot) Move(newLocation LatLong) {
    startingLocation := LatLong{0, 0}
    if r.Location != startingLocation { // Assume it takes no time to move to starting location
        time.Sleep(time.Duration(int64(r.Location.DistanceFrom(newLocation)) * int64(time.Hour) / speed))
    }
    r.Location = newLocation
}

func (r *Robot) IsNearTheTube(tubeStations []TubeStation) bool {
    return true
}
