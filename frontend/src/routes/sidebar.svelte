<script>
    import { RiSwordFill, RiGitRepositoryFill } from 'svelte-remixicon';
    let { handleNavigate } = $props();

</script>

<aside>
    <nav>
        <ul>
            <li>
                <button
                    class="icon"
                    type="button"
                    onclick={() => handleNavigate("Encounter")}
                    aria-label="Navigate to Encounter"
                >
                    <RiSwordFill size=32/>
                </button>
            </li>
            <li>
                <button
                    class="icon"
                    type="button"
                    onclick={() => handleNavigate("Library")}
                    aria-label="Navigate to Library"
                >
                    <RiGitRepositoryFill size=32/>
                </button>
            </li>
        </ul>
    </nav>
</aside>
<style>
    .icon {
        cursor: pointer;
        background: none;
        border: none;
        padding: 0;
        margin: 0;
        outline: none;
    }

    aside {
        width: 5%;
        background: #f4f4f4;
        height: 100%;
        padding-top: 10px;
        position: fixed;
        left: 0;
    }
    nav ul {
        list-style: none;
        padding: 0;
        margin: 0;
    }
    nav li {
        margin-bottom: 1em;
    }
</style>