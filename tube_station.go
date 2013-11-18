package main

import (
    "encoding/csv"
    "os"
    "strconv"
)

type TubeStation struct {
    Description string
    Location LatLong
}

func LoadTubeStations() []TubeStation {
    file, _ := os.Open("tube_stations.csv")
    defer file.Close()

    reader := csv.NewReader(file)
    records, _ := reader.ReadAll();

    tubeStations := make([]TubeStation, len(records))
    for i, fields := range(records) {
        lat, _ := strconv.ParseFloat(fields[1], 64)
        long, _ := strconv.ParseFloat(fields[2], 64)
        tubeStations[i] = TubeStation{fields[0], LatLong{Lat: lat, Long: long}}
    }

    return tubeStations
}
