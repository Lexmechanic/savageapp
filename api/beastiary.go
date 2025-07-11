package api

import (
	"encoding/json"
	"net/http"

	database "github.com/yourusername/savageapp/db"
)

func deleteBestiary(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	result, err := db.Exec("DELETE FROM bestiary WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete bestiary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No bestiary found with the given ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func getBestiary(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var bestiaryEntry database.Bestiary
	err := db.QueryRow("SELECT * FROM bestiary WHERE id = ?", r.FormValue("id")).Scan(
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
		http.Error(w, "Failed to query bestiary", http.StatusInternalServerError)
		return
	}

	// Convert the bestiary entry to JSON format
	bestiaryJson := bestiaryEntry.ConvertBestiaryToJson()

	skillRows, err := db.Query("SELECT bestiary_skills.bestiary_id, bestiary_skills.skill_id, bestiary_skills.rank, skills.name, skills.description, skills.default_skill FROM bestiary_skills INNER JOIN skills ON bestiary_skills.skill_id=skills.id WHERE bestiary_skills.bestiary_id = ?",
		r.FormValue("id"))
	if err != nil {
		http.Error(w, "Failed to query bestiary_skills", http.StatusInternalServerError)
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
			http.Error(w, "Failed to scan skill row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		skills = append(skills, skill)
	}

	if err := skillRows.Err(); err != nil {
		http.Error(w, "Error reading skills: "+err.Error(), http.StatusInternalServerError)
		return
	}

	bestiaryJson.Skills = skills
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestiaryJson)
}

func searchBestiary(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	playerCharacter := r.FormValue("player_character")
	if playerCharacter != "0" && playerCharacter != "1" {
		http.Error(w, "Invalid player_character value", http.StatusBadRequest)
		return
	}
	rows, err := db.Query("SELECT * FROM bestiary WHERE player_character = ? ORDER BY name ASC", playerCharacter)
	if err != nil {
		println(err.Error())
		http.Error(w, "Failed to query bestiary:"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bestiaryList := []database.BestiaryJson{}
	for rows.Next() {
		var bestiaryEntry database.Bestiary
		rows.ColumnTypes()
		err := rows.Scan(
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
			http.Error(w, "Failed to scan row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		bestiaryList = append(bestiaryList, bestiaryEntry.ConvertBestiaryToJson())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestiaryList)
}

func createBestiary(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var bestiaryEntry database.BestiaryJson
	if err := json.NewDecoder(r.Body).Decode(&bestiaryEntry); err != nil {
		println(err.Error())
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
		return
	}

	actionsJson, err := json.Marshal(bestiaryEntry.Actions)
	actionsString := `{"actions":` + string(actionsJson) + `}`
	if err != nil {
		http.Error(w, "Failed to marshal actions:"+err.Error(), http.StatusInternalServerError)
		return
	}

	traitsJson, err := json.Marshal(bestiaryEntry.Traits)
	traitsString := `{"traits":` + string(traitsJson) + `}`
	if err != nil {
		http.Error(w, "Failed to marshal traits:"+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := db.Exec("INSERT INTO bestiary (name, agility, smarts, spirit, strength, vigor, pace, toughness, parry, wounds, wounds_max, size, description, player_character, actions, traits) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", bestiaryEntry.Name, bestiaryEntry.Agility, bestiaryEntry.Smarts, bestiaryEntry.Spirit, bestiaryEntry.Strength, bestiaryEntry.Vigor, bestiaryEntry.Pace, bestiaryEntry.Toughness, bestiaryEntry.Parry, bestiaryEntry.Wounds, bestiaryEntry.WoundsMax, bestiaryEntry.Size, bestiaryEntry.Description, bestiaryEntry.PlayerCharacter, actionsString, traitsString)
	if err != nil {
		http.Error(w, "Failed to insert bestiary:"+err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	bestiaryEntry.ID = int(id)

	// Insert skills if any
	if len(bestiaryEntry.Skills) > 0 {
		for _, skill := range bestiaryEntry.Skills {
			_, err := db.Exec("INSERT INTO bestiary_skills (bestiary_id, skill_id, rank) VALUES (?, ?, ?)", bestiaryEntry.ID, skill.SkillId, skill.Rank)
			if err != nil {
				http.Error(w, "Failed to insert bestiary skill: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestiaryEntry)
}

func updateBestiary(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var bestiaryEntryJson database.BestiaryJson
	if err := json.NewDecoder(r.Body).Decode(&bestiaryEntryJson); err != nil {
		println(err.Error())
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
		return
	}

	actionsJson, err := json.Marshal(bestiaryEntryJson.Actions)
	actionsString := `{"actions":` + string(actionsJson) + `}`
	if err != nil {
		http.Error(w, "Failed to marshal actions:"+err.Error(), http.StatusInternalServerError)
		return
	}

	traitsJson, err := json.Marshal(bestiaryEntryJson.Traits)
	traitsString := `{"traits":` + string(traitsJson) + `}`
	if err != nil {
		http.Error(w, "Failed to marshal traits:"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE bestiary SET name = ?, agility = ?, smarts = ?, spirit = ?, strength = ?, vigor =?, pace = ?, toughness = ?, parry = ?, wounds = ?, wounds_max = ?, size = ?, description = ?, actions = ?, traits = ? WHERE id = ?",
		bestiaryEntryJson.Name,
		bestiaryEntryJson.Agility,
		bestiaryEntryJson.Smarts,
		bestiaryEntryJson.Spirit,
		bestiaryEntryJson.Strength,
		bestiaryEntryJson.Vigor,
		bestiaryEntryJson.Pace,
		bestiaryEntryJson.Toughness,
		bestiaryEntryJson.Parry,
		bestiaryEntryJson.Wounds,
		bestiaryEntryJson.WoundsMax,
		bestiaryEntryJson.Size,
		bestiaryEntryJson.Description,
		actionsString,
		traitsString,
		bestiaryEntryJson.ID,
	)
	if err != nil {
		http.Error(w, "Failed to insert bestiary:"+err.Error(), http.StatusInternalServerError)
		return
	}

	// First, delete existing skills for this bestiary
	_, err = db.Exec("DELETE FROM bestiary_skills WHERE bestiary_id = ?", bestiaryEntryJson.ID)
	if err != nil {
		http.Error(w, "Failed to delete existing skills: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Then, insert the new skills
	for _, skill := range bestiaryEntryJson.Skills {
		_, err := db.Exec("INSERT INTO bestiary_skills (bestiary_id, skill_id, rank) VALUES (?, ?, ?)", bestiaryEntryJson.ID, skill.SkillId, skill.Rank)
		if err != nil {
			http.Error(w, "Failed to insert bestiary skill: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	skillRows, err := db.Query("SELECT bestiary_skills.bestiary_id, bestiary_skills.skill_id, bestiary_skills.rank, skills.name, skills.description, skills.default_skill FROM bestiary_skills INNER JOIN skills ON bestiary_skills.skill_id=skills.id WHERE bestiary_skills.bestiary_id = ?",
		bestiaryEntryJson.ID)
	if err != nil {
		http.Error(w, "Failed to query bestiary_skills: "+err.Error(), http.StatusInternalServerError)
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
			http.Error(w, "Failed to scan skill row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		skills = append(skills, skill)
	}
	if err := skillRows.Err(); err != nil {
		http.Error(w, "Error reading skills: "+err.Error(), http.StatusInternalServerError)
		return
	}
	bestiaryEntryJson.Skills = skills
	// Return the updated bestiary entry

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestiaryEntryJson)
}

func BestiaryAPIRoutes() {
	http.HandleFunc("/api/bestiary/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createBestiary(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/bestiary/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			updateBestiary(w, r)
		case http.MethodDelete:
			deleteBestiary(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/bestiary", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			idValue := r.FormValue("id")
			if idValue != "" {
				getBestiary(w, r)
			} else {
				searchBestiary(w, r)
			}
		case http.MethodDelete:
			deleteBestiary(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
