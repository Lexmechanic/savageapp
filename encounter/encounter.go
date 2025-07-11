package encounter

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"

	database "github.com/yourusername/savageapp/db"
)

func (e *Encounter) StartEncounter() {
	e.Started = true
	// Wipe all held cards reshuffle a fresh deck and deal to all participants
	e.Deck.Suffle()
	for i, _ := range e.BestiaryList {
		e.BestiaryList[i].Cards = []int{}
		e.performDraw(i)
	}
	e.SortOrderByCard()
}

func (e *Encounter) EndEncounter() {
	e.Started = false
	e.Round = 0
	// Wipe all held cards
	for i := range e.BestiaryList {
		e.BestiaryList[i].Cards = []int{}
	}
}

func (e *Encounter) NextRound() {
	e.Round++
	for i := range e.BestiaryList {
		if !e.BestiaryList[i].Hold {
			e.BestiaryList[i].Cards = []int{}
			e.performDraw(i)
		}
	}
	e.SortOrderByCard()
}

func (e *Encounter) SortOrderByCard() {
	sort.SliceStable(e.BestiaryList, func(i, j int) bool {
		left := 0
		right := 0
		// If the left entry has cards, get the first card value
		if len(e.BestiaryList[i].Cards) > 0 {
			left = e.BestiaryList[i].Cards[0]

		}
		// If hold is true, add 52 to the card value
		if e.BestiaryList[i].Hold {
			left = left + 55
		}
		// If the right entry has cards, get the first card value
		if len(e.BestiaryList[j].Cards) > 0 {
			right = e.BestiaryList[j].Cards[0]
		}
		// If hold is true, add 52 to the card value
		if e.BestiaryList[j].Hold {
			right = right + 55
		}

		return left > right // DESC order
	})
}

// TODO: Add all the weird ass double draws take the lowest/mutlidraw take highest stuff
func (e *Encounter) performDraw(index int) {
	e.BestiaryList[index].Cards = append(e.BestiaryList[index].Cards, e.Deck.Draw())
	sort.Sort(sort.Reverse(sort.IntSlice(e.BestiaryList[index].Cards)))
}

func (e *Encounter) performDrawEncounterId(bestiaryListId int) {
	for index, bestiaryEntry := range e.BestiaryList {
		if bestiaryEntry.Id == bestiaryListId {
			e.BestiaryList[index].Cards = append(e.BestiaryList[index].Cards, e.Deck.Draw())
			sort.Sort(sort.Reverse(sort.IntSlice(e.BestiaryList[index].Cards)))
			e.SortOrderByCard()
			break
		}
	}

}

func (e *Encounter) SelectCard(bestiaryListId int, cardId int) {
	for index, bestiaryEntry := range e.BestiaryList {
		if bestiaryEntry.Id == bestiaryListId {
			temp := e.BestiaryList[index].Cards[cardId]
			e.BestiaryList[index].Cards[cardId] = e.BestiaryList[index].Cards[0]
			e.BestiaryList[index].Cards[0] = temp
			e.SortOrderByCard()
			break
		}
	}
}

func (e *Encounter) AddBestiary(beastiaryId int) {
	db := database.GetDB()
	var bestiaryEntry database.Bestiary
	err := db.QueryRow("SELECT * FROM bestiary WHERE id = ?", beastiaryId).Scan(
		&bestiaryEntry.ID,
		&bestiaryEntry.Name,
		&bestiaryEntry.Agility,
		&bestiaryEntry.Smarts,
		&bestiaryEntry.Spirit,
		&bestiaryEntry.Strength,
		&bestiaryEntry.Vigor,
		&bestiaryEntry.Pace,
		&bestiaryEntry.Toughness,
		&bestiaryEntry.Parry,
		&bestiaryEntry.Wounds,
		&bestiaryEntry.WoundsMax,
		&bestiaryEntry.Size,
		&bestiaryEntry.Description,
		&bestiaryEntry.Image,
		&bestiaryEntry.PlayerCharacter,
		&bestiaryEntry.Actions,
		&bestiaryEntry.Traits,
	)
	if err != nil {
		fmt.Println("Failed to query bestiary", err.Error())
		return
	}

	newEntry := EncounterBestiary{
		Id:            e.EntryCounter,
		BestiaryEntry: bestiaryEntry.ConvertBestiaryToJson(),
		Hidden:        false,
		Shaken:        false,
		Cards:         []int{},
		Hold:          false,
	}
	e.EntryCounter++

	newEntry.BestiaryEntry.Name = checkForDoubles(bestiaryEntry.Name, e)

	// If combat has started then add bestiary with the hidden attribute
	if e.Started {
		newEntry.Hidden = true
	}

	skillRows, err := db.Query("SELECT bestiary_skills.bestiary_id, bestiary_skills.skill_id, bestiary_skills.rank, skills.name, skills.description, skills.default_skill FROM bestiary_skills INNER JOIN skills ON bestiary_skills.skill_id=skills.id WHERE bestiary_skills.bestiary_id = ?",
		newEntry.BestiaryEntry.ID)
	if err != nil {
		fmt.Println("Error finding skills", err.Error())
		return
	}
	defer skillRows.Close()

	skills := []database.BestiarySkillsJson{}

	for skillRows.Next() {
		var skill database.BestiarySkillsJson
		err := skillRows.Scan(
			&skill.BestiaryId,
			&skill.SkillId,
			&skill.Rank,
			&skill.Name,
			&skill.Description,
			&skill.DefaultSkill,
		)
		if err != nil {
			fmt.Println("Error scanning skill rows", err.Error())
			return
		}
		skills = append(skills, skill)
	}

	newEntry.BestiaryEntry.Skills = skills

	e.BestiaryList = append(e.BestiaryList, newEntry)
}

func (e *Encounter) DeleteBestiary(id int) {
	for i, encounterEntry := range e.BestiaryList {
		if encounterEntry.Id == id {
			e.BestiaryList = slices.Delete(e.BestiaryList, i, i+1)
			break
		}
	}
}

func (e *Encounter) UpdateBestiary(updatedEntry EncounterBestiary) {
	for i, entry := range e.BestiaryList {
		if entry.Id == updatedEntry.Id {
			e.BestiaryList[i].Hidden = updatedEntry.Hidden
			e.BestiaryList[i].Shaken = updatedEntry.Shaken
			e.BestiaryList[i].Hold = updatedEntry.Hold
			e.BestiaryList[i].BestiaryEntry.Wounds = updatedEntry.BestiaryEntry.Wounds
			e.BestiaryList[i].BestiaryEntry.Name = updatedEntry.BestiaryEntry.Name
			break
		}
	}
}

func checkForDoubles(name string, e *Encounter) string {
	for _, encounterEntry := range e.BestiaryList {
		if encounterEntry.BestiaryEntry.Name == name {
			name = IncrementEndingNumber(name)
			name = checkForDoubles(name, e)
		}
	}

	return name
}

func IncrementEndingNumber(s string) string {
	re := regexp.MustCompile(`(\d+)$`)
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			inc := num + 1
			// Replace the last number with the incremented value
			return re.ReplaceAllString(s, " "+strconv.Itoa(inc))
		}
	}
	return s + " 1"
}
