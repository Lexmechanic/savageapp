<script>
    import { RiSkullLine, RiUserLine, RiGamepadLine, RiSpeedFill } from 'svelte-remixicon';
    import { onMount } from "svelte";
    import ListItem from "./list_item.svelte";
    import { searchBestiary } from '$lib/api';
    
    /**
     * @typedef {Object} Props
     * @property {boolean} [combat_started]
     * @property {any} addBestiary
     * @property {any} toggleCombat
     * @property {any} round
     * @property {any} nextRound
     */

    /** @type {Props} */
    let {
        combat_started = false,
        addBestiary,
        toggleCombat,
        round,
        nextRound
    } = $props();

    let bestiary = $state([]);
    let loading = $state(true);
    let error = $state(null);
    let activeTab = $state("control");
    let filter = $state("");

    async function fetchBeastiary() {
        loading = true;
        error = null;
        try {
            bestiary = await searchBestiary(activeTab);
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }

    function selectTab(tab) {
        activeTab = tab;
        filter = "";
        if (tab === "player") fetchBeastiary();
        if (tab === "bestiary") fetchBeastiary();
    }

    onMount(() => {
        fetchBeastiary();
    });

    // Filtered lists
    let filteredBeastiary = $derived(bestiary.filter(b => b.name.toLowerCase().includes(filter.toLowerCase())));
</script>
<div class="library_sidebar">
    <div class="tabs">
        <button class:active={activeTab === "control"} onclick={() => selectTab("control")}><RiGamepadLine size=25/></button>            
        <button class:active={activeTab === "player"} onclick={() => selectTab("player")}><RiUserLine size=25/></button>
        <button class:active={activeTab === "bestiary"} onclick={() => selectTab("bestiary")}><RiSkullLine size=25/></button>
    </div>
    {#if activeTab==="player" || activeTab === "bestiary"}
        <input
            class="filter-input"
            type="text"
            placeholder="Search..."
            bind:value={filter}
            autocomplete="off"
        />
    {/if}
    {#if loading}
        <p>Loading...</p>
    {:else if error}
        <p style="color: red;">{error}</p>
    {:else}
        {#if activeTab === "player" || activeTab === "bestiary"}
            <ul>
                {#each filteredBeastiary as bestiaryEntry, i}
                    <li class={i % 2 === 0 ? 'even' : 'odd'}>
                        <ListItem name={bestiaryEntry.name} id={bestiaryEntry.id} onClick={addBestiary}/>
                    </li>
                {/each}
            </ul>
        {:else if activeTab === "control"}
        <div class="control_tab">
            <button onclick={() => toggleCombat()} class={combat_started ? "combat_end_button" : "combat_start_button"}> {combat_started ? "End Combat" : "Start Combat"}</button>
            {#if combat_started == true}
                <table style="width: 100%;">
                    <thead>
                        <tr>
                            <th style="font-size: 1.5em;">Round: {round}</th>
                            <th style="text-align: right;"><button class="next_round_button" onclick={() => nextRound()}><RiSpeedFill size=25/></button></th>
                        </tr>
                    </thead>
                </table>
            {/if}
        </div>        
        {/if}
    {/if}
</div>
<style>
.control_tab {
    box-sizing: border-box;
    padding-left: 1em;
    padding-right: 1em;
    width: 100%; 
}

.combat_end_button {
    box-sizing: border-box;
    background: #bdbdbd;
    color: #000000;
    border: none;
    border-radius: 6px;
    padding-top: 1em;
    padding-bottom: 1em;
    font-size: 1em;
    font-weight: bold;
    cursor: pointer;
    width: 100%;
}

.combat_start_button {
    box-sizing: border-box;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding-top: 1em;
    padding-bottom: 1em;
    font-size: 1em;
    font-weight: bold;
    cursor: pointer;
    width: 100%;
}

.next_round_button {
    box-sizing: border-box;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 0.5em 0.5em 0.4em 0.5em;
    cursor: pointer;
    position: relative;
    top: 5px;
}

.combat_start_button:hover {
    background: #005fa3;
}

.combat_end_button:hover {
    background: #666666;    
}

.next_round_button:hover {
    background: #005fa3;
}

.library_sidebar {
    width: 15%;
    background: #f4f4f4;
    text-align: left;
    border-right: 1px solid #242323;
    border-left: 1px solid #242323;
    height: 100%;
    position: fixed;
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
    width: 90%;
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

.library_sidebar li.odd:hover, .library_sidebar li.even:hover{
    background: #ffffff;
}

</style>