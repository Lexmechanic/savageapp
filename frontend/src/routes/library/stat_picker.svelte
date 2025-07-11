<script>
    import { onMount } from "svelte";
    import { RiDeleteBinLine } from 'svelte-remixicon';
    /**
     * @typedef {Object} Props
     * @property {number} [value]
     * @property {any} selected
     * @property {any} stat
     * @property {number} [id]
     * @property {boolean} [deleteEnabled]
     * @property {any} deleteSkill
     */

    /** @type {Props} */
    let {
        value = $bindable(0),
        selected = $bindable(),
        stat = $bindable(),
        id = 0,
        deleteEnabled = false,
        deleteSkill
    } = $props();

    function selectDie(die) {
        selected = die;
        stat = `${selected}+${value}`;
    }

    function onValueInput(e) {
        value = e.target.value;
        stat = `${selected}+${value}`;
    }

    onMount(() => {
        if (stat) {
            const [die, val] = stat.split("+");
            selected = die || "";
            value = val || "";
        }
    });
</script>

    <div class="stat-picker">
        <div class="dice-buttons">
            <button class:selected={selected === "d4"} onclick={() => selectDie("d4")}>d4</button>
            <button class:selected={selected === "d6"} onclick={() => selectDie("d6")}>d6</button>
            <button class:selected={selected === "d8"} onclick={() => selectDie("d8")}>d8</button>
            <button class:selected={selected === "d10"} onclick={() => selectDie("d10")}>d10</button>
            <button class:selected={selected === "d12"} onclick={() => selectDie("d12")}>d12</button>
        </div>
        <span class="plus-sign">+</span>
        <input
            type="number"
            bind:value
            class="stat-input"
            placeholder="Value..."
            oninput={onValueInput}
        />
        {#if deleteEnabled}
            <button class="delete-button" onclick={() =>deleteSkill(id)}><RiDeleteBinLine/></button>
        {/if}
    </div>

<style>
.delete-button {
    background: none;
    border: none;
    cursor: pointer;
    color: #3d3d3d;
    border-radius: 4px;
    border: 1px solid #888;
    text-align: center;
}

.delete-button:hover {
    background: #ff0000;
}

.stat-picker {
    display: flex;
    align-items: center;
    gap: .3em;
}

.plus-sign {
    font-size: 2em;
    font-weight: bold;
    color: #444;
    position: relative;
    top: -6px;
}

.dice-buttons {
    display: flex;
    gap: 0.3em;
}

.dice-buttons button {
    padding: 0.4em 0.7em;
    border: 1px solid #888;
    background: #f4f4f4;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
    transition: background 0.2s, border 0.2s;
}

.dice-buttons button.selected,
.dice-buttons button:focus {
    background: #007acc;
    color: #fff;
    border-color: #007acc;
}

.stat-input {
    width: 60px;
    padding: 0.3em;
    border: 1px solid #bbb;
    border-radius: 4px;
    font-size: 1em;
}
</style>