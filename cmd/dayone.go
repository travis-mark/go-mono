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