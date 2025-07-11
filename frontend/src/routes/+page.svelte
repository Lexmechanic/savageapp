<script>
    import Sidebar from "./sidebar.svelte";
    import Encounter from "./encounter/encounter.svelte";
    import Library from "./library/library.svelte";

    let mode = $state("none")
    let current = $state("Encounter");

    function handleNavigate(name) {
        current = name;
    }

    function selectMode(selectedMode) {
        mode = selectedMode;
    }
</script>

<main>
    {#if mode === "none"}
        <h1>Welcome to the Savage App</h1>
        <p>Select your role.</p>
        <div class="center-buttons">
            <button onclick={() => selectMode("GM")}>Game Master</button>
            <button onclick={() => selectMode("Player")}>Player</button>
        </div>
    {:else if mode === "GM"}
        <Sidebar handleNavigate={handleNavigate}/>
        {#if current === "Encounter"}
            <Encounter mode = {mode} />
        {:else if current === "Library"}
            <Library />
        {/if}
    {:else if mode === "Player"}
        <Encounter mode = {mode} />
    {/if}    
</main>

<style>
	main {
        padding: 0;
		text-align: center;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

</style>