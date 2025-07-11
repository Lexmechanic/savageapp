<script>
    import { RiHeart3Fill, RiHeart3Line, RiEyeLine, RiEyeOffLine, RiHeartPulseFill, RiHeartPulseLine, RiHand, RiFileCopyFill } from 'svelte-remixicon';
    import WoundPopup from "./wound_popup.svelte";
    import Card from "./card.svelte";
    let {
        encounter_bestiary = $bindable(),
        index,
        selected = $bindable(),
        updateEncouterEntry,
        drawCardForEntry,
        selectCard,
        mode
    } = $props();
    let show_wound_popup = $state(false);
    let temp

    function showWoundPopup() {
        if(mode === "GM") {
           show_wound_popup = true; 
        } 
    }

    function toggleHidden() {
        if(mode != "GM"){return}
        if(encounter_bestiary.hidden == false){
            encounter_bestiary.hidden = true;
        } else {
            encounter_bestiary.hidden = false;
        }
        updateEncouterEntry(encounter_bestiary);
    }

    function toggleShaken() {
        if(mode != "GM"){return}
        if(encounter_bestiary.shaken == false){
            encounter_bestiary.shaken = true;
        } else {
            encounter_bestiary.shaken = false;
        }
        updateEncouterEntry(encounter_bestiary);
    }

    // Toggle the hold state of the encounter_bestiary
    function toggleHold() {
        if(mode != "GM" && encounter_bestiary.bestiary_entry.player_character == false){return}
        if(encounter_bestiary.hold == false){
            encounter_bestiary.hold = true;
        } else {
            encounter_bestiary.hold = false;
        }
        updateEncouterEntry(encounter_bestiary);        
    }

    function woundsOrMax(){
        if(encounter_bestiary.bestiary_entry.wounds <= encounter_bestiary.bestiary_entry.wounds_max){
            return encounter_bestiary.bestiary_entry.wounds_max
        } 
        return encounter_bestiary.bestiary_entry.wounds
    }

</script>

<tr class={index % 2 === 0 ? 'odd' : 'even'} onclick={() => selected = encounter_bestiary}>
    <th>
        <button type="button" class="icon_button" onclick={() => toggleHidden()}>
            {#if mode == 'GM' && encounter_bestiary.hidden == false} 
                <RiEyeLine size=20/>
            {:else if mode == 'GM'}
                <RiEyeOffLine size=20/>
            {/if}
        </button>
        {encounter_bestiary.bestiary_entry.name}
    </th>
    <th>
        <div class="vertical-buttons">
            {#if mode === "GM" || encounter_bestiary.bestiary_entry.player_character == true || encounter_bestiary.hold == true}
            <button 
                class={encounter_bestiary.hold ? "action_button_down" : "action_button"}
                title="Hold toggle"
                onclick={() => toggleHold() }
            >
                <RiHand size=20/>
            </button>
            {/if}
            {#if mode === "GM" || encounter_bestiary.bestiary_entry.player_character == true}
                <button class="action_button" title="Draw another card" onclick={() => drawCardForEntry(encounter_bestiary.id)}><RiFileCopyFill size=20/></button>
            {/if}
        </div>
    </th>
    <th class="card-row">
        {#each encounter_bestiary.cards as card, card_index}
            <button type="button" class="icon_button" onclick={() => selectCard(encounter_bestiary.id, card_index)}><Card card={card}/></button>
        {/each}
    </th>
    <th>
        {#if mode === "GM" || encounter_bestiary.bestiary_entry.player_character == true}
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <button type="button" class={mode === "GM" ? "icon_button" : "icon_button_player"} onclick={() => showWoundPopup(encounter_bestiary)}>
                {#each Array(woundsOrMax()) as _, i}
                    {#if encounter_bestiary.bestiary_entry.wounds_max <= i}
                        <RiHeart3Fill size=20 color="gold"/>
                    {:else if encounter_bestiary.bestiary_entry.wounds > i}
                        <RiHeart3Fill size=20 color="#ff3232"/>
                    {:else}
                        <RiHeart3Line size=20/>
                    {/if}
                {/each}
            </button>
        {/if}
        <button type="button" class={mode === "GM" ? "icon_button" : "icon_button_player"} onclick={() => toggleShaken()}>
            {#if encounter_bestiary.shaken == true}
                <RiHeartPulseFill size=20/>
            {:else if mode=="GM"}
                <RiHeartPulseLine size=20/>
            {/if}
        </button>
    </th>
    <th>{encounter_bestiary.bestiary_entry.toughness}</th>
    <th>{encounter_bestiary.bestiary_entry.parry}</th>
</tr>
{#if show_wound_popup}
    <WoundPopup 
    bind:encounter_bestiary={encounter_bestiary}
    updateEncouterEntry={updateEncouterEntry}
    on:close={() => {show_wound_popup = false}} />
{/if}

<style>
tr.even {
    background: #f9f9f9;
}

tr.odd {
    background: #e6e6e6;
}

th {
    padding-top: 0.5em;
    padding-bottom: .5em;
    padding-left: .25em;
    text-align: left;
}

.card-row {
    display: flex;
    flex-direction: row;
    gap: 0.5em;
    align-items: center;
}

.vertical-buttons {
    display: flex;
    flex-direction: column;
    gap: 0.3em;
    align-items: flex-start;
}

.icon_button {
    cursor: pointer;
    background: none;
    border: none;
    padding: 0;
    margin: 0;
    display: inline-flex;
    align-items: center;
}

.icon_button_player {
    background: none;
    border: none;
    padding: 0;
    margin: 0;
    display: inline-flex;
    align-items: center;
}

.action_button {
    background: #007acc;
    color: white;
    border: none;
    padding: 0.5em .5em;
    border-radius: 4px;
    cursor: pointer;
}

.action_button_down {
    background: #bdbdbd;
    color: black;
    border: none;
    padding: 0.5em .5em;
    border-radius: 4px;
    cursor: pointer;
}


</style>

