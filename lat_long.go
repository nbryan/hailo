package main

import (
    "math"
)

const earthsRadius float64 = 6371000.0 // Meters

type LatLong struct {
    Lat float64
    Long float64
}

func (a LatLong) DistanceFrom(b LatLong) float64 {
    // Spherical law of cosines formula
    return math.Acos(math.Sin(radians(a.Lat)) * math.Sin(radians(b.Lat)) + math.Cos(radians(a.Lat)) * math.Cos(radians(b.Lat)) * math.Cos(radians(b.Long) - radians(a.Long))) * earthsRadius
}

func radians(degrees float64) float64 {
    return degrees * math.Pi / 180
}
