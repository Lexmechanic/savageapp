<script>
    import { run } from 'svelte/legacy';
    import {getSkillById, deleteSkill, updateSkill, createSkill} from '$lib/api.js';
    import { onMount } from "svelte";
    import { RiDeleteBinLine } from 'svelte-remixicon';

    /**
     * @typedef {Object} Props
     * @property {any} [id]
     * @property {any} [skill]
     * @property {any} refreshList
     */

    /** @type {Props} */
    let { id = $bindable(null), skill = $bindable({}), refreshList } = $props();
    let create = $state(false);
    let stat_types = ['Agility', 'Smarts', 'Spirit', 'Strength', 'Vigor'];
    let loading = $state(false);
    let error = $state(null);

    async function fetchSkillEntry() {
        loading = true;
        error = null;
        create = false;
        try {
            skill = await getSkillById(id);
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }

    async function newSkillEntry() {
        create = true;
        skill = {
            name: "New Skill",
            description: "none",
            skill_stat: "Agility",
            default_skill: false,
        };
    }

    onMount(() => {
        if (!id) {
            newSkillEntry();
        } else {
            fetchSkillEntry();
        }
    });

    // Remount/fetch when id changes
    run(() => {
        if (id) {
            fetchSkillEntry();
        } else if (!id){
            newSkillEntry();
        }
    });

    async function saveSkillEntry() {
        try {
            loading = true;
            if (create === true) {
                skill = await createSkill(skill)
            } else {
                skill = await updateSkill(skill)
            }    
        } catch (e) {
            error = e.message;
            alert("Error updating skill entry: " + error);
        } finally {
            create = false;              
            loading = false;
            refreshList(skill.id);            
        }

    }

    async function deleteSkillEntry() {
        try {
            const res = await deleteSkill(id);
        } catch(e) {
            alert("Error deleting skill entry: " + e.message);
        } finally {
            refreshList();
        }
        id = null;
    }

</script>
{#if loading}
    <p>Loading skill entry...</p>
{:else if error}
    <p style="color:red;">{error}</p>
{:else}
    <div class="skill-details">
            <table>
                <thead>
                    <tr>
                        <th><label for="skill-name-input">Skill Name:</label></th>
                        <th class="text-input">
                            <input
                                id="skill-name-input"
                                type="text"
                                bind:value={skill.name}
                                class="skill-name-input"
                                placeholder="Player Name"
                            />
                        </th>
                    </tr>
                    <tr>
                        <th><label for="skill-description-input">Description:</label></th>
                        <th class="text-input">
                            <textarea
                                    id="skill-description-input"
                                    type="text"
                                    bind:value={skill.description}
                                    class="skill-description-input"
                                    placeholder="Player Description"
                                    rows="4"
></textarea>
                        </th>    
                    </tr>
                    <tr>
                        <th><label for="skill-stat-input">Skill Stat:</label></th>
                        <th class="skill-details-table">
                            <select class="player-skill-select" bind:value={skill.skill_stat} id="skill-stat-input">
                                {#each stat_types as stat}
                                    <option value={stat}>{stat}</option>
                                {/each}
                            </select>
                        </th>
                    </tr>
                    <tr>
                        <th><label for="skill-default-input">Default Skill:</label></th>
                        <th class="skill-details-table">
                            <input
                                type="checkbox"
                                id="skill-default-input"
                                bind:checked={skill.default_skill}
                                class="skill-default-checkbox"
                            />
                        </th>
                    </tr>
                </thead>    
            </table>
    </div>
    <div class="bottom-save-div">
        {#if !create}<button class="delete-btn" onclick={deleteSkillEntry}><RiDeleteBinLine/></button>{/if}
        <button class="save-btn" onclick={saveSkillEntry}>Save</button>
    </div>
{/if}

<style>
.skill-default-checkbox {
    margin-left: 0;
    margin-right: 0.5em;
    vertical-align: middle;
}

.skill-details {
    position:fixed;
    left: 25%;
    width: 75%;
    height: 100%;
    overflow-y: scroll; /* Enable vertical scrolling */
    padding-left: 1em;
    padding-right: 1em;
}

.skill-name-input {
    font-size: 1.1em;
    padding: 0.2em 0.5em;
    margin-bottom: 0.7em;
    border: 1px solid #bbb;
    border-radius: 4px;
    width: 100%;
}

.skill-description-input {
    font-size: 1.1em;
    padding: 0.2em 0.5em;
    margin-bottom: 0.7em;
    border: 1px solid #bbb;
    border-radius: 4px;
    resize: none;
    width: 100%;
}

.text-input {
    width: 100%;
    padding-right: 2em;
}

th {
    text-align: left;
}

.save-btn {
    position: fixed;
    right: .25em;
    bottom: .25em;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 1em 3em;
    font-size: 0.9em;
    font-weight: bold;
    cursor: pointer;
    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
    transition: background 0.2s;
}
.save-btn:hover {
    background: #005fa3;
}

.bottom-save-div {
min-height: fit-content;
}

.delete-btn {
    position: fixed;
    right: 9em;
    bottom: .25em;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 1em 1em;
    font-size: 0.9em;
    font-weight: bold;
    cursor: pointer;
    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
    transition: background 0.2s;
}
.delete-btn:hover {
    background: #ff0000;
}

label {
    width: 120px;
    margin-top: -1em;
}
</style>   