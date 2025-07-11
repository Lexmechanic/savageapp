package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Bestiary struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Agility         string         `json:"agility"`
	Smarts          string         `json:"smarts"`
	Spirit          string         `json:"spirit"`
	Strength        string         `json:"strength"`
	Vigor           string         `json:"vigor"`
	Pace            int            `json:"pace"`
	Toughness       string         `json:"toughness"`
	Parry           int            `json:"parry"`
	Wounds          int            `json:"wounds"`
	WoundsMax       int            `json:"wounds_max"`
	Size            int            `json:"size"`
	Image           sql.NullString `json:"image"`
	Description     sql.NullString `json:"description"`
	PlayerCharacter int            `json:"player_character"`
	Actions         string         `json:"actions"`
	Traits          string         `json:"traits"`
}

type BestiaryJson struct {
	ID              int                  `json:"id"`
	Name            string               `json:"name"`
	Agility         string               `json:"agility"`
	Smarts          string               `json:"smarts"`
	Spirit          string               `json:"spirit"`
	Strength        string               `json:"strength"`
	Vigor           string               `json:"vigor"`
	Pace            int                  `json:"pace"`
	Toughness       string               `json:"toughness"`
	Parry           int                  `json:"parry"`
	Wounds          int                  `json:"wounds"`
	WoundsMax       int                  `json:"wounds_max"`
	Size            int                  `json:"size"`
	Description     string               `json:"description"`
	PlayerCharacter int                  `json:"player_character"`
	Skills          []BestiarySkillsJson `json:"skills"`
	Actions         []Action             `json:"actions"`
	Traits          []Trait              `json:"traits"`
}

type Action struct {
	Damage      string `json:"damage"`
	Description string `json:"description"`
}

type Trait struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Beastiary_Skills struct {
	BestiaryId int    `json:"bestiary_id"`
	SkillId    string `json:"id"`
	Rank       string `json:"rank"`
}

type BestiarySkillsJson struct {
	BestiaryId   int    `json:"bestiary_id"`
	SkillId      int    `json:"skill_id"`
	Name         string `json:"name"`
	Rank         string `json:"rank"`
	Description  string `json:"description"`
	DefaultSkill bool   `json:"default_skill"`
}

func (bestiary *Bestiary) ConvertBestiaryToJson() BestiaryJson {
	bestiaryJson := BestiaryJson{
		ID:              bestiary.ID,
		Name:            bestiary.Name,
		Agility:         bestiary.Agility,
		Smarts:          bestiary.Smarts,
		Spirit:          bestiary.Spirit,
		Strength:        bestiary.Strength,
		Vigor:           bestiary.Vigor,
		Pace:            bestiary.Pace,
		Toughness:       bestiary.Toughness,
		Parry:           bestiary.Parry,
		Wounds:          bestiary.Wounds,
		WoundsMax:       bestiary.WoundsMax,
		Size:            bestiary.Size,
		Description:     bestiary.Description.String,
		PlayerCharacter: bestiary.PlayerCharacter,
	}

	actionBytes := []byte(bestiary.Actions)
	err := json.Unmarshal(actionBytes, &bestiaryJson)
	if err != nil {
		fmt.Println("Error unmarshalling actions", (bestiary.Actions), ":", err)
	}

	traitBytes := []byte(bestiary.Traits)
	err = json.Unmarshal(traitBytes, &bestiaryJson)
	if err != nil {
		fmt.Println("Error unmarshalling traits", (bestiary.Traits), ":", err)
	}
	return bestiaryJson
}
