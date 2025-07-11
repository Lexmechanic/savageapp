<script>
    /**
     * @typedef {Object} Props
     * @property {any} [encounter_bestiary]
     */

    /** @type {Props} */
    let { encounter_bestiary = null } = $props();
</script>
<div class="pane">
{#if encounter_bestiary === null}
    <h3>No Selection</h3>
{:else}
    <h3>{encounter_bestiary.bestiary_entry.name}</h3>
        <table class="stat_grid">
            <thead>
                <tr>
                    <th>Agility</th>
                    <th>Smarts</th>
                    <th>Spirit</th>
                </tr>
                <tr>
                    <td>{encounter_bestiary.bestiary_entry.agility}</td>
                    <td>{encounter_bestiary.bestiary_entry.smarts}</td>
                    <td>{encounter_bestiary.bestiary_entry.spirit}</td>
                </tr>
                <tr>
                    <th>Strength</th>
                    <th>Vigor</th>
                    <th>Pace</th>
                </tr>
                <tr>
                    <td>{encounter_bestiary.bestiary_entry.strength}</td>
                    <td>{encounter_bestiary.bestiary_entry.vigor}</td>
                    <td>{encounter_bestiary.bestiary_entry.pace}</td>
                </tr>
                <tr>
                    <th>Size</th>
                    <th>Tough</th>
                    <th>Parry</th>
                </tr>
                <tr>
                    <td>{encounter_bestiary.bestiary_entry.size}</td>
                    <td>{encounter_bestiary.bestiary_entry.toughness}</td>
                    <td>{encounter_bestiary.bestiary_entry.parry}</td>
                </tr>
            </thead>
        </table>
        <hr class="fade"/>
        <table class="skill_grid">
            <thead>
                <tr>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
            {#each encounter_bestiary.bestiary_entry.actions as action, i}
                <tr class={i % 2 === 0 ? 'odd' : 'even'}>
                    <td><span class="bold">{action.damage}:&nbsp</span>{action.description}</td>
                </tr>
            {/each}
            </tbody>
        </table>    
        <hr class="fade"/>
        <table class="skill_grid">
            <thead>
                <tr>
                    <th>Traits</th>
                </tr>
            </thead>
            <tbody>
            {#each encounter_bestiary.bestiary_entry.traits as trait, i}
                <tr class={i % 2 === 0 ? 'odd' : 'even'}>
                    <td><span class="bold">{trait.name}:&nbsp</span>{trait.description}</td>
                </tr>
            {/each}
            </tbody>
        </table>    
        <hr class="fade"/>
        <table class="skill_grid">
            <thead>
                <tr>
                    <th colspan="2">Skills</th>
                </tr>
            </thead>
            <tbody>
            {#each encounter_bestiary.bestiary_entry.skills as skill, i}
                <tr class={i % 2 === 0 ? 'odd' : 'even'}>
                    <td>{skill.name}:</td>
                    <td>{skill.rank}</td>
                </tr>
            {/each}
            </tbody>
        </table>    
        <hr class="fade"/>
        <table class="skill_grid">
            <thead>
                <tr>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>{encounter_bestiary.bestiary_entry.description}:</td>
                </tr>
            </tbody>
        </table>
        <p></p>    
{/if}
</div>

<style>
hr.fade {
    border: 0;
    height: 1px;
    background-image: linear-gradient(to right, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.75), rgba(0, 0, 0, 0));
}

table.stat_grid {
    width: 95%;
    margin-left: auto;
    margin-right: auto;
    border-radius: 6px;
    border: 1px solid #444;
    border-collapse: separate;   
    border-spacing: 0; 
}

.bold {
    font-weight: bold;
}

table.stat_grid td{
    width: 30%;
}

table.skill_grid {
    width: 95%;
    margin-left: auto;
    margin-right: auto;
    border-radius: 6px;
    border: 1px solid #444;
    border-collapse: separate;   
    border-spacing: 0; 
}

th, td {
    padding: 0.25em .1em;
}

th {
    background: #007acc;
    color:white;
}

.pane {
    text-align: center;
    background: #f4f4f4;
    height: 100%;
    overflow-y: auto;
}

tr.even {
    background: #f9f9f9;
}

tr.odd {
    background: #e6e6e6;
}
</style>