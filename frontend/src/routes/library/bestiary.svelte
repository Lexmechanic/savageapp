<script>
    import { run } from 'svelte/legacy';
    import {getBestiaryById, deleteBestiary, searchSkills, updateBestiary, createBestiary} from '$lib/api.js';
    import { onMount } from "svelte";
    import Stat_picker from "./stat_picker.svelte";
    import { RiDeleteBinLine } from 'svelte-remixicon';
    import AddSkillPopup from "./add_skill_popup.svelte"
    import EntryListItem from "./entry_list_item.svelte";

    /**
     * @typedef {Object} Props
     * @property {any} [id]
     * @property {any} [bestiaryEntry]
     * @property {any} refreshList
     * @property {any} type
     */

    /** @type {Props} */
    let {
        id = $bindable(null),
        bestiaryEntry = $bindable({}),
        refreshList,
        type
    } = $props();

    let create = $state(false);
    let loading = $state(false);
    let error = $state(null);
    let showSkillPopup = $state(false);
    let showActionPopup = $state(false);

    async function fetchBestiaryEntry() {
        loading = true;
        error = null;
        create = false;
        try {
            bestiaryEntry = await getBestiaryById(id);
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }

    function newTrait(name, description) {
        var newTraitEntry = {
            name: name,
            description: description
        }
        bestiaryEntry.traits = [...bestiaryEntry.traits, newTraitEntry];
    }

    function deleteTrait(index) {
        if (!bestiaryEntry.traits || bestiaryEntry.traits.length === 0) return;
        bestiaryEntry.traits = bestiaryEntry.traits.filter((_, i) => i !== index);
    }

    function newAction(damage, description) {
        var newActionEntry = {
            damage: damage,
            description: description
        }
        bestiaryEntry.actions = [...bestiaryEntry.actions, newActionEntry];
    }

    function deleteAction(index) {
        if (!bestiaryEntry.actions || bestiaryEntry.actions.length === 0) return;
        bestiaryEntry.actions = bestiaryEntry.actions.filter((_, i) => i !== index);
    }

    // Add a passed in skill database object to the bestiary entry
    function addSkill(skill) {
        var newSkillEntry = {
            skill_id: skill.id,
            name: skill.name,
            rank: "d4+0",
            description: skill.description
        }
        bestiaryEntry.skills = [...bestiaryEntry.skills, newSkillEntry]
    }

    function deleteSkillFromBestiaryEntry(skillId) {
        if (!bestiaryEntry.skills) return;
        bestiaryEntry.skills = bestiaryEntry.skills.filter(skill => skill.skill_id !== skillId);
    }

    async function newBestiaryEntry() {
        create = true;
        bestiaryEntry = {
            name: "New Player",
            description: "none",
            agility: "d4+0",
            smarts: "d4+0",
            spirit: "d4+0",
            strength: "d4+0",
            vigor: "d4+0",
            parry: 2,
            toughness: "2(0)",
            size: 0,
            pace: 6,
            wounds: 3,
            player_character: 1,
            actions: [],
            traits: []
        };
        if (type === "bestiary") {
            bestiaryEntry.player_character = 0;
            bestiaryEntry.name = "New Bestiary Entry";
        }
        var skills = [];
        try {
            skills = await searchSkills(true);
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
        bestiaryEntry.skills = skills.map(skill => ({
            skill_id: skill.id,
            name: skill.name,
            description: skill.description,
            default_skill: skill.default_skill,
            rank: "d4+0"
        }));
    }

    onMount(() => {
        if (!id) {
            newBestiaryEntry();
        } else {
            fetchBestiaryEntry();
        }
    });
    
    // Remount/fetch when id changes
    run(() => {
        if (id) {
            fetchBestiaryEntry();
        } else if (!id){
            newBestiaryEntry();
        }
    });

    async function deleteBestiaryEntry() {
        try {
            const res = await deleteBestiary(id); 
        } catch(e) {
            alert("Error deleting bestiary entry: " + e.message);
        } finally {
            refreshList();
        }
        id = null;
    }

    async function saveBestiaryEntry() {
        let url;
        bestiaryEntry.wounds_max = bestiaryEntry.wounds
        try {
            loading = true;
            error = null;
            if (create === true) {
                bestiaryEntry = await createBestiary(bestiaryEntry);
            } else {
                bestiaryEntry = await updateBestiary(bestiaryEntry);
            }
        } catch (e) {
            alert("Error updating bestiary entry: " + e.message);
        } finally {
            create = false;                  
            loading = false;
            refreshList(bestiaryEntry.id);
        }
    }
</script>

{#if loading}
    <p>Loading bestiary entry...</p>
{:else if error}
    <p style="color:red;">{error}</p>
{:else}
    <div class="bestiary-details">
        <table>
            <thead>
                <tr>
                    <th><label for="bestiary-name-input">{#if  type === "bestiary"}Entry Name:{:else }Player Name:{/if}</label></th>
                    <th class="text-input">
                        <input
                            id="bestiary-name-input"
                            type="text"
                            bind:value={bestiaryEntry.name}
                            class="bestiary-name-input"
                            placeholder="Player Name"
                        />
                    </th>
                </tr>
                <tr>
                    <th><label for="bestiary-description-input">Description:</label></th>
                    <th class="text-input">
                        <textarea
                                id="bestiary-description-input"
                                type="text"
                                bind:value={bestiaryEntry.description}
                                class="bestiary-description-input"
                                placeholder="Player Description"
                                rows="4"
                        ></textarea>
                    </th>    
                </tr>
            </thead> 
        </table>    
        <div class="bestiary-stats-grid">
            <div class="bestiary-stats">
                <h3>Player Stats</h3>
                <table class="stats-list">
                    <thead>
                        <tr class="stat-row">
                            <th>Agility:</th>
                            {#key bestiaryEntry.agility}
                            <th><Stat_picker bind:stat={bestiaryEntry.agility}/></th>
                            {/key}
                        </tr>  
                        <tr>
                            <th>Smarts:</th>
                            {#key bestiaryEntry.smarts}
                            <th><Stat_picker bind:stat={bestiaryEntry.smarts}/></th>
                            {/key}
                        </tr>
                        <tr>
                            <th>Spirit:</th>
                            {#key bestiaryEntry.spirit}
                            <th><Stat_picker bind:stat={bestiaryEntry.spirit}/></th>
                            {/key}
                        </tr>
                        <tr>
                            <th>Strength:</th>
                            {#key bestiaryEntry.strength}
                            <th><Stat_picker bind:stat={bestiaryEntry.strength}/></th>
                            {/key}
                        </tr>
                        <tr>
                            <th>Vigor:</th>
                            {#key bestiaryEntry.vigor}
                            <th><Stat_picker bind:stat={bestiaryEntry.vigor}/></th>
                            {/key}
                        </tr>
                    </thead>
                </table>
            </div>
            <div class="bestiary-stats">
                <h3>Derived Stats</h3>
                <table class="stats-list">
                    <thead>
                        <tr>
                            <th>Parry:</th>
                            <th>                
                                <input
                                    type="number"
                                    bind:value={bestiaryEntry.parry}
                                    class="bestiary-name-input"
                                    placeholder="Parry"
                                />
                            </th>
                        </tr>
                        <tr>
                            <th>Toughness:</th>
                            <th>                
                                <input
                                    type="text"
                                    bind:value={bestiaryEntry.toughness}
                                    class="bestiary-name-input"
                                    placeholder="Parry"
                                />
                            </th>
                        </tr>
                        <tr>
                            <th>Size:</th>
                            <th>                
                                <input
                                    type="number"
                                    bind:value={bestiaryEntry.size}
                                    class="bestiary-name-input"
                                    placeholder="Parry"
                                />
                            </th>
                        </tr>
                        <tr>
                            <th>Pace:</th>
                            <th>                
                                <input
                                    type="number"
                                    bind:value={bestiaryEntry.pace}
                                    class="bestiary-name-input"
                                    placeholder="Parry"
                                />
                            </th>
                        </tr>
                        <tr>
                            <th>Wounds:</th>
                            <th>                
                                <input
                                    type="number"
                                    bind:value={bestiaryEntry.wounds}
                                    class="bestiary-name-input"
                                    placeholder="Parry"
                                />
                            </th>
                        </tr>
                    </thead>
                </table>
            </div>
            <div class="bestiary-stats">
                <h3>Skills</h3>
                {#if bestiaryEntry.skills && bestiaryEntry.skills.length > 0}
                    <table class="stats-list">
                        <thead>                      
                        {#each bestiaryEntry.skills as skill}
                            <tr>
                                <th>{skill.name}:</th>
                                <th><Stat_picker bind:stat={skill.rank} id={skill.skill_id} deleteSkill={deleteSkillFromBestiaryEntry} deleteEnabled=true/></th>
                            </tr>
                        {/each}
                        </thead>                        
                    </table>

                {:else}
                    <p>No skills defined.</p>
                {/if}
                <div><button class="add-item-btn" onclick={() => {showSkillPopup = true}}>Add New Skill</button></div>                
            </div>
            <div class="bestiary-stats">
                <h3>Actions</h3>
                {#if bestiaryEntry.actions && bestiaryEntry.actions.length > 0}
                    {#each bestiaryEntry.actions as action, ia}
                        <div class="input-row">
                            <label for="action-name-input" class="trait-action-label">Damage:</label>
                            <input
                                style="width: 100%;"
                                id="action-name-input"
                                type="text"
                                placeholder="Ex. 1d6+0 or None"
                                autocomplete="off"
                                bind:value={action.damage}
                            >
                            <button class="delete-button" onclick={() => deleteAction(ia)}><RiDeleteBinLine/></button>
                        </div>
                        <div class="input-row">
                            <label for="action-desc-input" class="trait-action-label">Description:</label>
                            <textarea
                                style="width: 100%;"
                                id="action-desc-input"
                                placeholder="Action Description"
                                rows="4"
                                bind:value={action.description}
                            ></textarea>    
                        </div>
                    {/each}
                {:else}
                    <p>No actions defined.</p>
                {/if}
                <div><button class="add-item-btn" onclick={() => newAction("","")}>Add New Action</button></div>                
            </div>
            <div class="bestiary-stats">
                <h3>Traits</h3>
                {#if bestiaryEntry.traits && bestiaryEntry.traits.length > 0}
                    {#each bestiaryEntry.traits as trait, it}
                            <div class="input-row">
                                <label for="trait-name-input" class="trait-action-label">Name:</label>
                                <input
                                    style="width: 100%;"
                                    id="trait-name-input"
                                    type="text"
                                    placeholder="Trait Name"
                                    autocomplete="off"
                                    bind:value={trait.name}
                                >
                                <button class="delete-button" onclick={() => deleteTrait(it)}><RiDeleteBinLine/></button>
                            </div>
                            <div class="input-row">
                                <label for="trait-desc-input" class="trait-action-label">Description:</label>
                                <textarea
                                    style="width: 100%;"
                                    id="trait-desc-input"
                                    placeholder="Action Description"
                                    rows="4"
                                    bind:value={trait.description}
                                ></textarea>    
                            </div>
                        {/each}
                {:else}
                    <p>No traits defined.</p>
                {/if}
                <div><button class="add-item-btn" onclick={() => newTrait("","")}>Add New Trait</button></div>                
            </div>
        </div>
    </div>
    <div class="bottom-save-div">
        {#if !create}<button class="delete-btn" onclick={deleteBestiaryEntry}><RiDeleteBinLine/></button>{/if}
        <button class="save-btn" onclick={saveBestiaryEntry}>Save</button>
    </div>
    {#if showSkillPopup}
        <AddSkillPopup addSkill={addSkill} on:close={() => {showSkillPopup = false}} />
    {/if}
    {#if showActionPopup}
        <AddActionPopup addAction={addAction} on:close={() => {showActionPopup = false}} />
    {/if}
{/if}

<style>
.stats-list {
    width: 100%;
    border-collapse: collapse;
}

.input-row {
    display: flex;
    gap: 0.7em;
    font-weight: bold;
}

.bestiary-stats-grid {
    padding-left: 1em;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 2em;
    margin-top: 1.5em;
    width: 100%;
    max-width: 600px;
}

@media (max-width: 1500px) {
    .bestiary-stats-grid {
        grid-template-columns: 1fr;
        max-width: 600px;
    }
}

.bestiary-stats {
    border: black 1px solid;
    padding-left: 1em;
    padding-bottom: 1em;
    padding-right: 1em;
    width: 525px;
    box-sizing: border-box;
    background: #fafbfc;
    padding-bottom: 1em;
}

.bestiary-details {
    position:relative;
    left: 25%;
    width: 75%;
    overflow-y: scroll; /* Enable vertical scrolling */
    padding-bottom: 1em;
}

.bestiary-name-input {
    font-size: 1.1em;
    padding: 0.2em 0.5em;
    margin-bottom: 0.7em;
    border: 1px solid #bbb;
    border-radius: 4px;
    width: 100%;
}

.bestiary-description-input {
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

.add-item-btn {
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: .5em .5em;
    font-size: 0.9em;
    font-weight: bold;
    cursor: pointer;
    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
    transition: background 0.2s;
    width: 100%;
}

.add-item-btn:hover {
    background: #005fa3;
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

.bottom-save-div {
min-height: fit-content;
}

.trait-action-label {
    text-align: right;
    margin-top: .5em;
}
</style>