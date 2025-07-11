<script>
    import { RiBardFill } from 'svelte-remixicon';
    /**
     * @typedef {Object} Props
     * @property {any} id
     * @property {any} name
     * @property {any} type
     * @property {any} onEdit
     * @property {boolean} [starred]
     */

    /** @type {Props} */
    let {
        id,
        name,
        type,
        onEdit,
        starred = false
    } = $props();
</script>

<div class="list-item">
    <span>{#if starred}<RiBardFill />{/if}{name}</span>
    <button class="edit-btn" onclick={() => onEdit(id,type)}>Edit</button>
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

.list-item:hover .edit-btn,
.list-item:focus-within .edit-btn {
    opacity: 1;
    pointer-events: auto;
}
</style>