export const API_BASE = ""

// searchBestiary returns a list of bestiary entries based on the type specified.
// @type {string} = Type of bestiary you want to search for, e.g., "bestiary", "player", etc.
export async function searchBestiary(type) {
    let url = `${API_BASE}/api/bestiary`
            if (type !== "bestiary"){
                url = url + "?player_character=1";
            } else if (type !== "player"){
                url = url + "?player_character=0";                
            }
    const res = await fetch(url);
    if (!res.ok) throw new Error("Failed to fetch bestiary");
    return res.json();
}

// searchSkills returns a list of skills from the API.
// @param {boolean} only_default_skill - If true, only returns the default skills.
export async function searchSkills(only_default_skill = false) {
    const url = `${API_BASE}/api/skill${only_default_skill ? '?default_skill=true' : ''}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error("Failed to fetch skills");
    return res.json();
}

export async function getSkillById(id) {
    const url = `${API_BASE}/api/skill?id=${id}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error("Failed to fetch skill");
    return res.json();
}

export async function updateSkill(skill) {
    const url = `${API_BASE}/api/skill/update`;
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(skill)
    });
    if (!res.ok) throw new Error("Failed to update skill");
    return res.json();
}

export async function createSkill(skill) {
    const url = `${API_BASE}/api/skill/create`;
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(skill)
    });
    if (!res.ok) throw new Error("Failed to create skill");
    return res.json();
}

export async function deleteSkill(id) {
    const url = `${API_BASE}/api/skill?id=${id}`;
    const res = await fetch(url, {
        method: 'DELETE'
    });
    if (!res.ok) throw new Error("Failed to delete skill");
    return res.json();
}

// getBestiaryById fetches a single bestiary entry by its ID.
// @param {number} id - The ID of the bestiary entry to fetch.
export async function getBestiaryById(id) {
    const url = `${API_BASE}/api/bestiary?id=${id}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error("Failed to fetch bestiary entry");
    return res.json();
}

export async function deleteBestiary(id) {
    const url = `${API_BASE}/api/bestiary?id=${id}`;
    const res = await fetch(url, {
        method: 'DELETE'
    });
    console.log("Deleting bestiary entry with id: " + res);
    if (!res.ok) throw new Error("Failed to delete bestiary entry");
    return res.json();
}

export async function updateBestiary(bestiary) {
    const url = `${API_BASE}/api/bestiary/update`;
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(bestiary)
    });
    if (!res.ok) throw new Error("Failed to update bestiary entry");
    return res.json();
}

export async function createBestiary(bestiary) {
    const url = `${API_BASE}/api/bestiary/create`;
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(bestiary)
    });
    if (!res.ok) throw new Error("Failed to create bestiary entry");
    return res.json();
}