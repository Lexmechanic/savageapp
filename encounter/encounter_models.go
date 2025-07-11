package encounter

import (
	"sync"

	database "github.com/yourusername/savageapp/db"
)

var (
	encounter *Encounter
	once      sync.Once
)

type Encounter struct {
	BestiaryList []EncounterBestiary `json:"bestiary_list"`
	Deck         Deck                `json:"deck"`
	Round        int                 `json:"round"`
	Started      bool                `json:"started"`
	EntryCounter int                 `json:"entry_counter"`
}

type EncounterBestiary struct {
	Id            int                   `json:"id"`
	Hidden        bool                  `json:"hidden"`
	Cards         []int                 `json:"cards"`
	BestiaryEntry database.BestiaryJson `json:"bestiary_entry"`
	Shaken        bool                  `json:"shaken"`
	Hold          bool                  `json:"hold"`
}

func GetEncounter() *Encounter {
	once.Do(func() {
		encounter = &Encounter{
			Round:        0,
			BestiaryList: []EncounterBestiary{},
			Deck:         NewDeck(),
			Started:      false,
			EntryCounter: 0,
		}
		go handleBroadCast()
	})
	return encounter
}
