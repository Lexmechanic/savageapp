<script>
    import { RiDeleteBinLine } from 'svelte-remixicon';
    let {
        name,
        description,
        index,
        deleteAction
        } = $props();
</script>

<tr>
    <td style="text-align: left;"><span style="font-weight: bold;padding-right: .5em;">{name}: </span><span style="padding-right: .5em;">{description}</span></td>
    <td style="text-align: right;"><button class="delete-button" onclick={() => deleteAction(i)}><RiDeleteBinLine/></button></td>
</tr>

<style>
.delete-button {
    background: none;
    border: none;
    cursor: pointer;
    color: #3d3d3d;
    border-radius: 4px;
    border: 1px solid #888;
    text-align: right;
}

.delete-button:hover {
    background: #ff0000;
}
</style>