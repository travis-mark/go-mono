package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Journal struct {
	Metadata JournalMetadata `json:"metadata"`
	Entries    []JournalEntry `json:"entries"`
}

type JournalMetadata struct {
    Version string `json:"version"`
}

type JournalEntry struct {
    IsAllDay bool `json:"isAllDay"`
    Starred bool `json:"starred"`
    Text string `json:"text"`
    Location JournalLocation `json:"location"`
    TimeZone string `json:"timeZone"`
    Duration float64 `json:"duration"`
    Uuid string `json:"uuid"`
    IsPinned bool `json:"isPinned"`
    ModifiedDate string `json:"modifiedDate"`
    Weather JournalWeather `json:"weather"`
    CreationDate string `json:"creationDate"`

}

type JournalLocation struct {
    Region JournalRegion `json:"region"`
    Country string `json:"country"`
    Longitude float64 `json:"longitude"`
    AdministrativeArea string `json:"administrativeArea"`
    PlaceName string `json:"placeName"`
    Latitude float64 `json:"latitude"`
}

type JournalRegion struct {
    Center JournalCenter `json:"center"`
    Radius float64 `json:"radius"`
}

type JournalCenter struct {
    Longitude float64 `json:"longitude"`
    Latitude float64 `json:"latitude"`
}

type JournalWeather struct {
    SunsetDate string `json:"sunsetDate"`
    TemperatureCelsius float64 `json:"temperatureCelsius"`
    WeatherServiceName string `json:"weatherServiceName"`
    WindBearing float64 `json:"windBearing"`
    SunriseDate string `json:"sunriseDate"`
    ConditionsDescription string `json:"conditionsDescription"`
    PressureMB float64 `json:"pressureMB"`
    VisibilityKM float64 `json:"visibilityKM"`
    RelativeHumidity float64 `json:"relativeHumidity"`
    WindSpeedKPH float64 `json:"windSpeedKPH"`
    WeatherCode string `json:"weatherCode"`
    WindChillCelsius float64 `json:"windChillCelsius"`
}

func main() {
    path := "/Users/travis/Desktop/04-13-2025_13-20-/Journal.json"
    file, err := os.Open(path)
    if err != nil {
        log.Fatal("open file failed with error: ", err)
    }
    defer file.Close()

    var journal Journal
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&journal); err != nil {
        log.Fatal("decode failed with error: ", err)
    }

    if len(journal.Entries) == 0 {
        fmt.Println("WARNING: File ihas no entries")
        return
    }

    fmt.Println("First entry in the JSON file:")
	fmt.Printf("%+v\n", journal.Entries[0])
}