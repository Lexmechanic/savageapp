<script>
    import { onMount, onDestroy } from "svelte";
    import { RiShieldCrossLine, RiSwordLine } from 'svelte-remixicon';
    import { page } from '$app/state';
    import EncounterControl from "./encounter_control.svelte"
    import EncounterListItem from "./encounter_list_item.svelte";
    import EncounterBestiaryInfo from "./encounter_bestiary_info.svelte";

    let ws;
    let encounter = $state({bestiary_list: [], started: false});
    let message;
    let selected_encounter_bestiary = $state(null);
    const API_WEBSOCKET = '/ws'
    let { mode } = $props();

    function setupWebsocket() {
        ws = new WebSocket(API_WEBSOCKET);
        ws.onopen = () => {
            getEncounter()
        };
        ws.onmessage = (event) => {
            try {
                encounter = JSON.parse(event.data);
            } catch (e) {
                console.error("Invalid JSON from server:", event.data);
            }
        };
        ws.onerror = (err) => {
            console.error("WebSocket error:", err);
        };
        ws.onclose = () => {
            console.log("WebSocket closed");
            setTimeout(setupWebsocket, 1000)
        }; 
    }

    onMount(() => {
        setupWebsocket()
        window.addEventListener("keydown", handleKeydown);

    });

    onDestroy(() => {
        if (ws) ws.close();
        window.removeEventListener("keydown", handleKeydown);
    });

    function handleKeydown(event) {
        if (event.key === "Delete" && selected_encounter_bestiary && mode === "GM") {
            message = {
                action: "delete-bestiary",
                id: selected_encounter_bestiary.id
            };
            if (ws && ws.readyState === WebSocket.OPEN) {
                ws.send(JSON.stringify(message));
            }
            selected_encounter_bestiary = null            
        } 
    }

    function addBestiary(id) {
        message = {
            action: "add-bestiary",
            id: id
        };
        if (ws && ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify(message));
        }
    }

    function getEncounter() {
        message = {
                action: "get-encounter"
            };
        if (ws && ws.readyState === WebSocket.OPEN) {        
            ws.send(JSON.stringify(message));
        }
    }

    function toggleCombat() {
        if(encounter.started === false) {
            message = {action: "start-encounter"};    
        } else {
             message = {action: "end-encounter"};           
        }
        if (ws && ws.readyState === WebSocket.OPEN) {        
            ws.send(JSON.stringify(message));
        }    
    }

    function nextRound() {
        if(encounter.started === true) {
            message = {action: "next-round"}; 
            if (ws && ws.readyState === WebSocket.OPEN) {        
                ws.send(JSON.stringify(message));
            }      
        } 
    }

    function updateEncouterEntry(entry) {
        message = {
            action: "update-bestiary",
            encounter_bestiary: entry
        };
        if (ws && ws.readyState === WebSocket.OPEN) {        
            ws.send(JSON.stringify(message));
        }
    }

    function drawCardForEntry(id) {
        message = {
            action: "draw-card",
            id: id
        };
        if (ws && ws.readyState === WebSocket.OPEN) {        
            ws.send(JSON.stringify(message));
        }
    }

    function selectCard(id, cardId) {
        message = {
            action: "select-card",
            id: id,
            card_id: cardId
        };
        if (ws && ws.readyState === WebSocket.OPEN) {        
            ws.send(JSON.stringify(message));
        }
    }
    
    let visibleEncounterList = $derived(encounter.bestiary_list
        .filter(entry =>
            mode === "GM" ||
            (entry.bestiary_entry.player_character == true &&
            entry.hidden == false) ||
            (encounter.started == true &&
            entry.hidden == false)
        ));
</script>

<div class={mode === "GM" ? 'encounter_area' : 'encounter_area_player'}>
    <table class="encounter_table">
        <thead>
            <tr class="encounter_table_header">
                <th>Name</th>
                <th></th>
                <th>Draws</th>
                <th>Wounds</th>
                <th><RiShieldCrossLine/></th>
                <th><RiSwordLine/></th>
            </tr>
        </thead>
        {#each visibleEncounterList as encounter_entry, i}
            <EncounterListItem 
                bind:encounter_bestiary={visibleEncounterList[i]}
                index={i}
                bind:selected={selected_encounter_bestiary}
                updateEncouterEntry={updateEncouterEntry}
                mode={mode}
                drawCardForEntry={drawCardForEntry}
                selectCard={selectCard}
                />
        {/each}
    </table>    
</div>
{#if mode === "GM"} 
    <div class="control_area">
            <EncounterControl 
            addBestiary={addBestiary} 
            toggleCombat={toggleCombat} 
            combat_started={encounter.started} 
            round={encounter.round}
            nextRound={nextRound}
            />
            
    </div>
    <div class="right_pane">
        <EncounterBestiaryInfo encounter_bestiary={selected_encounter_bestiary}/>
    </div>
{/if}

<style>

.control_area {
    z-index: 1000;
}

.right_pane {
    width: 15%;
    overflow-y: auto;
    height: 100%;
    position: absolute;
    right: 0%;
    top: 0%;    
    border-left: 1px solid #242323;
}

.encounter_area {
    left: 20%;
    width: 65%;
    overflow-y: auto;
    position: absolute;
    max-height: 100%;
}

.encounter_area_player {
    width: 100%;
    overflow-y: auto;
    max-height: 100%;
    z-index: -1;
}

.encounter_table {
    border-collapse: separate;
    border-spacing: 0px;
    width: 100%;
    padding-bottom: 1em;
}

.encounter_table_header {
    background: #007acc;
    color:white;

}

th {
    padding-top: 0.75em;
    padding-bottom: .75em;
    padding-left: .25em;
    text-align: left;    
}
</style>