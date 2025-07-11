<script>
    import { RiSearchLine } from 'svelte-remixicon';
    let { id, name, onClick } = $props();
    let showPopup = $state(false);
    let mouseX = $state(0);
    let mouseY = $state(0);

    function handleMouseMove(event) {
        mouseX = event.clientX;
        mouseY = event.clientY;
    }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="list-item" onclick={() => onClick(id)}>
    <span>{name}</span>
    <button class="edit-btn" 
        onmouseenter={() => showPopup = true}
        onmouseleave={() => showPopup = false}
        onmousemove={handleMouseMove}
    >
        <RiSearchLine />
    </button>
    {#if showPopup}
        <div
            class="popup-tip"
            style="left: {mouseX + 10}px; top: {mouseY + 10}px; position: fixed;"
        >
            This is a popup!
        </div>
    {/if}
</div>

<style>
.list-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5em 0.5em;
    border-bottom: 1px solid #ccc;
    position: relative;
}

.edit-btn {
    background: #007acc;
    color: #fff;
    border: none;
    border-radius: 3px;
    padding: 0.25em 0.75em;
    cursor: pointer;
    font-size: 0.9em;
    transition: background 0.2s;
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.2s;
}

.popup-tip {
    background: #222;
    color: #fff;
    padding: 0.5em 1em;
    border-radius: 6px;
    font-size: 0.95em;
    white-space: nowrap;
    box-shadow: 0 2px 8px rgba(0,0,0,0.18);
    pointer-events: none;
}

.list-item:hover .edit-btn,
.list-item:focus-within .edit-btn {
    opacity: 1;
    pointer-events: auto;
}
</style>