<script>
    import { RiSkullLine, RiUserLine, RiQuillPenAiLine } from 'svelte-remixicon';
    import {searchSkills, searchBestiary} from '$lib/api.js';
    import { onMount } from "svelte";
    import ListItem from "./list_item.svelte";
    import Beastiary from "./bestiary.svelte";
    import Skill from "./skill.svelte";

    let bestiary = $state([]);
    let skills = $state([]);
    let loading = $state(true);
    let error = $state(null);
    let activeTab = $state("player");
    let filter = $state("");
    let selectedType = $state(null);
    let selectedID = $state(null);

    async function fetchBeastiary(id = null) {
        loading = true;
        error = null;
        try {
            bestiary = await searchBestiary(activeTab);
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
        selectedID = id;
    }

    async function fetchSkills(id = null) {
        loading = true;
        error = null;
        try {
            skills = await searchSkills();
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
        selectedID = id;
    }

    function selectTab(tab) {
        activeTab = tab;
        filter = "";
        if (tab === "player") fetchBeastiary();
        if (tab === "bestiary") fetchBeastiary();
        if (tab === "skill") fetchSkills();
    }

    function onEdit(id, type) {
        selectedType = type;
        selectedID = id;
    }


    onMount(() => {
        fetchBeastiary();
    });

    // Filtered lists
    let filteredBeastiary = $derived(bestiary.filter(b => b.name.toLowerCase().includes(filter.toLowerCase())));
    let filteredSkills = $derived(skills.filter(s => s.name.toLowerCase().includes(filter.toLowerCase())));
</script>
<div>
    <div class="library_sidebar">
        <div class="tabs">
            <button class:active={activeTab === "player"} onclick={() => selectTab("player")}><RiUserLine size={25}/></button>
            <button class:active={activeTab === "bestiary"} onclick={() => selectTab("bestiary")}><RiSkullLine size={25}/></button>
            <button class:active={activeTab === "skill"} onclick={() => selectTab("skill")}><RiQuillPenAiLine size={25}/></button>
        </div>
        <input
            class="filter-input"
            type="text"
            placeholder="Search..."
            bind:value={filter}
            autocomplete="off"
        />
        <button
            class="add-btn"
            title="Add New"
            onclick={() => onEdit(null, activeTab)}
        >+</button>
        {#if loading}
            <p>Loading...</p>
        {:else if error}
            <p style="color: red;">{error}</p>
        {:else}
            {#if activeTab === "player" || activeTab === "bestiary"}
                <ul>
                    {#each filteredBeastiary as bestiaryEntry, i}
                        <li class={i % 2 === 0 ? 'even' : 'odd'}>
                            <ListItem name={bestiaryEntry.name} id={bestiaryEntry.id} onEdit={onEdit} type={activeTab} />
                        </li>
                    {/each}
                </ul>
            {:else if activeTab === "skill"}
                <ul>
                    {#each filteredSkills as skill, i}
                        <li class={i % 2 === 0 ? 'even' : 'odd'}>
                            <ListItem name={skill.name} starred={skill.default_skill} id={skill.id} onEdit={onEdit} type={activeTab} />
                        </li>
                    {/each}
                </ul>
            {/if}
        {/if}
    </div>
    <div>
        {#if selectedType === "player"}
            <Beastiary id={selectedID} type={selectedType} refreshList={fetchBeastiary}/>
        {:else if selectedType === "bestiary"}
            <Beastiary id={selectedID} type={selectedType} refreshList={fetchBeastiary}/>
        {:else if selectedType === "skill"}
            <Skill id={selectedID} type={selectedType} refreshList={fetchSkills}/>
        {/if}
    </div>
</div>
<style>
.library_sidebar {
    width: 20%;
    background: #f4f4f4;
    text-align: left;
    border-right: 1px solid #242323;
    border-left: 1px solid #242323;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 5%;
    overflow-y: auto;
}

.tabs {
    display: flex;
    justify-content: space-between;
    margin-bottom: 1em;
}

.tabs button {
    flex: 1;
    padding: 0.5em;
    background: #e0e0e0;
    border: none;
    border-bottom: 2px solid transparent;
    cursor: pointer;
    font-weight: bold;
    transition: background 0.2s, border-bottom 0.2s;
}

.tabs button.active {
    background: #fff;
    border-bottom: 2px solid #007acc;
}

.filter-input {
    width: 80%;
    margin: 0.5em 0 1em 0;
    padding: 0.4em;
    border: 1px solid #bbb;
    border-radius: 4px;
    font-size: 1em;
    box-sizing: border-box;
}

.library_sidebar ul {
    list-style: none;
    padding-left: 0;
    margin: 0;
}

.library_sidebar li.even {
    background: #f9f9f9;
}

.library_sidebar li.odd {
    background: #e6e6e6;
}

.add-btn {
    padding: 0.4em 1em;
    font-size: 1em;
    font-weight: bold;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background 0.2s;
    margin-left: 0.25em;
    flex: 0 0 auto;
}
.add-btn:hover {
    background: #005fa3;
}
</style>