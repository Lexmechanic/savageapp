package db

import "database/sql"

type Skill struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	DefaultSkill bool           `json:"default_skill"`
	SkillStat    string         `json:"skill_stat"`
}

type SkillJson struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DefaultSkill bool   `json:"default_skill"`
	SkillStat    string `json:"skill_stat"`
}

func ConvertSkillToJson(skill Skill) SkillJson {
	return SkillJson{
		ID:           skill.ID,
		Name:         skill.Name,
		Description:  skill.Description.String,
		DefaultSkill: skill.DefaultSkill,
		SkillStat:    skill.SkillStat,
	}
}
