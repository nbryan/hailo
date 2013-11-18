package main

import (
    "fmt"
)

func main() {
    tubeStations := LoadTubeStations()

    for _, station := range tubeStations {
        fmt.Println(station)
    }
}
