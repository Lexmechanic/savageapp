CREATE TABLE IF NOT EXISTS bestiary (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    agility TEXT NOT NULL,
    smarts TEXT NOT NULL,
    spirit TEXT NOT NULL,
    strength TEXT NOT NULL,
    vigor TEXT NOT NULL,
    pace INTEGER NOT NULL,
    toughness INTEGER NOT NULL,
    parry INTEGER NOT NULL,
    wounds INTEGER NOT NULL,
    wounds_max INTEGER NOT NULL,
    size INTEGER DEFAULT 0,
    description TEXT DEFAULT "none",
    image BINARY,
    player_character BOOLEAN DEFAULT 0,
    actions TEXT DEFAULT '{"actions": []}',
    traits TEXT DEFAULT '{"traits": []}'
);

CREATE TABLE IF NOT EXISTS skills (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    skill_stat TEXT NOT NULL,
    description TEXT DEFAULT "none",
    default_skill BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS bestiary_skills (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bestiary_id INTEGER NOT NULL,
    skill_id INTEGER NOT NULL,
    rank TEXT NOT NULL DEFAULT "d4+0",
    FOREIGN KEY (bestiary_id) REFERENCES bestiary(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE ON UPDATE CASCADE
);