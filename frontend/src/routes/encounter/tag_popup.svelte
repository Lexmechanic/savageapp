<script>
    import { self, createBubbler, preventDefault } from 'svelte/legacy';
    import { onMount } from 'svelte';

    const bubble = createBubbler();
    import { RiCheckLine } from 'svelte-remixicon';
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
    let { 
        addTag
    } = $props();
    
    let tagInput = $state("");
    let inputElement;

    onMount(() => {
        if (inputElement) {
            inputElement.focus();
            inputElement.select();
        }
    });

    function confirm() {
        addTag(tagInput);
        close();
    }

    function close() {
        dispatch("close");
    }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="popup-backdrop" onclick={close}>
    <div class="popup-content" onclick={(event) => event.stopPropagation()}>
        <input
            bind:this={inputElement}
            type="text"
            bind:value={tagInput}
            placeholder="Tag Text..."
        />
        <button class="close-btn" onclick={confirm} title="Confirm"><RiCheckLine size=25/></button>        
    </div>
</div>

<style>
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
    background: #ffffff;
    border-radius: 8px;
    padding: 2em;
    box-shadow: 0 8px 32px rgba(0,0,0,0.18);
    width: 33%;
    height:fit-content ;
    position: relative;
    text-align: left;
    align-items: left;
    font-weight: bold;
    color: #444;
}

.stat-input {
    width: 60px;
    padding: 0.3em;
    border: 1px solid #bbb;
    border-radius: 4px;
    font-size: 1em;
}

.close-btn {
    box-sizing: border-box;
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    position: absolute;
    right: 1em;
}
.close-btn:hover {
    background: #005fa3;
}

</style>