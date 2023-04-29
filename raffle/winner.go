package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const path = "raffle.json"

// raffleEntry is the struct we unmarshal raffle entries into
type RaffleEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []RaffleEntry {
	data, err := ioutil.ReadFile(path)

	var raffleEntries []RaffleEntry

	content := string(data)
	err = json.Unmarshal([]byte(content), &raffleEntries)
	if err != nil {
		panic(err)
	}

	return raffleEntries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []RaffleEntry) RaffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
