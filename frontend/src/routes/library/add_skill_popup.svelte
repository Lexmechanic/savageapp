<script>
    import { self } from 'svelte/legacy';
    import {searchSkills} from '$lib/api.js';

    import { createEventDispatcher, onMount } from "svelte";
    const dispatch = createEventDispatcher();
    let skills = $state([])
    let error
    let filter = $state("");
    let { addSkill } = $props();

    function close() {
        dispatch("close");
    }

    function selectItem(skill) {
        addSkill(skill)
        close()
    }

    onMount(() => {
        fetchSkills()
    });

    async function fetchSkills() {
        error = null;
        try {
            skills = await searchSkills();
        } catch (e) {
            error = e.message;
        }
    }

    // Filtered lists
    let filteredSkills = $derived(skills.filter(s => s.name.toLowerCase().includes(filter.toLowerCase())));

</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="popup-backdrop" onclick={self(close)}>
    <div class="popup-content" onclick={(event) => event.stopPropagation()}>
        <div>
            <input
                class="filter-input"
                type="text"
                placeholder="Search..."
                bind:value={filter}
                autocomplete="off"
            />
        </div>
            <table class="skill-table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Attribute</th>
                        <th>Description</th>  
                    </tr>
                </thead>
                <tbody>
                    {#each filteredSkills as skill, i}
                    <tr class={i % 2 === 0 ? 'even' : 'odd'} onclick={() => selectItem(skill)}>
                        <td style="width:15%">{skill.name}</td>
                        <td style="width:10%">{skill.skill_stat}</td>
                        <td style="width:75%">{skill.description}</td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        <button class="close-btn" onclick={close} title="Close">&times;</button>             
        </div>
</div>

<style>
.skill-table tr.even {
    background: #f9f9f9;
}

.skill-table tr.odd {
    background: #e6e6e6;
}

.skill-table tr:hover {
    background: #007acc;
    color: #fff;
}

.skill-table th {
    background: #007acc;
    color: #fff;
}

.popup-backdrop {
    position: fixed;
    top: 0; left: 0; right: 0; bottom: 0;
    background: rgba(0,0,0,0.35);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.popup-content {
    background: #fff;
    border-radius: 8px;
    padding: 4em 2.5em 1.5em 2.5em;
    box-shadow: 0 8px 32px rgba(0,0,0,0.18);
    width: 70%;
    height: 50%;
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    overflow-y: auto;
}

.skill-table {
    width: 100%;
    border: 1px solid #888;
}

.close-btn {
    position: absolute;
    top: 0em;
    right: 0em;
    background: none;
    border: none;
    font-size: 2em;
    color: #888;
    cursor: pointer;
    z-index: 10;
    transition: color 0.2s;
}
.close-btn:hover {
    color: #c00;
}

.filter-input {
    width: 300px;
    margin: 0.5em 0 1em 0;
    padding: 0.4em;
    border: 1px solid #bbb;
    border-radius: 4px;
    font-size: 1em;
    position: absolute;
    top: 1em;
    left: 2.5em;
}

</style>