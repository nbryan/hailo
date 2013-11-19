package main

import (
    "encoding/csv"
    "os"
    "fmt"
    "strconv"
    "time"
)

type Instruction struct {
    Location LatLong
    Time time.Time
}

var tubeStations []TubeStation // Evil global variable

func main() {
    tubeStations = LoadTubeStations()

    ids := []int{5937, 6043}
    done := make(chan int)

    for _, id := range ids {
        go Dispatch(Robot{id, LatLong{0, 0}}, done)
    }

    for i := len(ids); i > 0; {
        <-done
        i--
    }
}

func Dispatch(r Robot, done chan int) {
    fmt.Println("Dispatch")
    ch := make(chan Instruction, 10) // Only allow 10 instructions in queue at a time
    go ReadInstructions(r.ID, ch)

    for instruction := range ch {
        r.ReceiveInstruction(instruction)
        fmt.Println(r.ID, instruction)
    }

    done <- 1
}

func ReadInstructions(id int, ch chan Instruction) {
    fmt.Println("ReadInstructions")
    file, _ := os.Open(fmt.Sprintf("%v.csv", id))
    defer file.Close()

    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err != nil {
            break
        }

        lat, _ := strconv.ParseFloat(record[1], 64)
        long, _ := strconv.ParseFloat(record[2], 64)
        time, _ := time.Parse("2006-01-02 15:04:05", record[3])
        instruction := Instruction{LatLong{lat, long}, time}

        ch <- instruction
    }

    close(ch)
}
