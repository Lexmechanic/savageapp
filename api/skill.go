package api

import (
	"encoding/json"
	"net/http"

	database "github.com/yourusername/savageapp/db"
)

func createSkill(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var skill database.SkillJson
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO skills (name, skill_stat, description, default_skill) VALUES (?, ?, ?, ?)",
		skill.Name, skill.SkillStat, skill.Description, skill.DefaultSkill)
	if err != nil {
		http.Error(w, "Failed to create skill: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	skill.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skill)
}

func searchSkill(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	defaultWhere := ""
	if r.URL.Query().Get("default_skill") != "" {
		defaultWhere = " WHERE default_skill =" + r.URL.Query().Get("default_skill")
	}

	rows, err := db.Query("SELECT id, name, skill_stat, description, default_skill FROM skills " + defaultWhere + " ORDER BY default_skill DESC, name ASC")
	if err != nil {
		http.Error(w, "Failed to search skills: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	skills := []database.SkillJson{}
	for rows.Next() {
		var skill database.Skill
		if err := rows.Scan(&skill.ID, &skill.Name, &skill.SkillStat, &skill.Description, &skill.DefaultSkill); err != nil {
			http.Error(w, "Failed to scan skill: "+err.Error(), http.StatusInternalServerError)
			return
		}
		skills = append(skills, database.ConvertSkillToJson(skill))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skills)
}

func getSkill(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	var skill database.Skill
	err := db.QueryRow("SELECT id, name, skill_stat, description, default_skill FROM skills WHERE id = ?", id).Scan(
		&skill.ID, &skill.Name, &skill.SkillStat, &skill.Description, &skill.DefaultSkill)
	if err != nil {
		http.Error(w, "Failed to scan row", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.ConvertSkillToJson(skill))
}

func updateSkill(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var skill database.SkillJson
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE skills SET name = ?, skill_stat = ?, description = ?, default_skill = ? WHERE id = ?",
		skill.Name, skill.SkillStat, skill.Description, skill.DefaultSkill, skill.ID)
	if err != nil {
		http.Error(w, "Failed to update skill: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skill)
}

func deleteSkill(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	result, err := db.Exec("DELETE FROM skills WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete skill: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No skill found with the given ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

// Starts all the http handers related to skills
func SkillAPIRoutes() {
	http.HandleFunc("/api/skill/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createSkill(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/skill/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			updateSkill(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/skill", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			idValue := r.FormValue("id")
			if idValue != "" {
				getSkill(w, r)
			} else {
				searchSkill(w, r)
			}
		case http.MethodDelete:
			deleteSkill(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
