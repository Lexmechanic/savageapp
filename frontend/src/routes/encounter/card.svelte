<script>
    const suits = ["♣", "♦", "♥", "♠"];
    const values = ["2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "☻"];
    let { card } = $props();

    function getSuit(cardNumber) {
        let partial = (cardNumber/4 - Math.floor(cardNumber/4))*4
        return partial
    }
</script>
        <div class="card-rectangle">
            {#if getSuit(card) == 0 || getSuit(card) == 3}
                <div class="card-value">{values[Math.floor(card/4)]}</div>           
                <div class="card-suit-top">{suits[getSuit(card)]}</div>
                <div class="card-suit-bottom">{suits[getSuit(card)]}</div>
            {:else}
                <div class="card-value-red">{values[Math.floor(card/4)]}</div>
                <div class="card-suit-top-red">{suits[getSuit(card)]}</div>
                <div class="card-suit-bottom-red">{suits[getSuit(card)]}</div>
            {/if}
        </div>
<style>
.card-rectangle {
    width: 63px;
    height: 88px;
    background: #fff;
    border: 2px solid #333;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.12);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', Arial, sans-serif;
    font-size: 2em;
    position: relative;
}

.card-value {
    position: absolute;
    top: 0.5em;
    font-size: 1.2em;
    font-weight: bold;
    text-align: center;
    top: 16px
}

.card-value-red {
    position: absolute;
    top: 0.5em;
    font-size: 1.2em;
    font-weight: bold;
    text-align: center;
    top: 16px;
    color: red;
}

.card-suit-top {
    font-size: 0.75em;
    position: absolute;
    left: 3px;
    top: -3px;
}

.card-suit-bottom {
    font-size: 0.75em;
    position: absolute;
    right: 3px;
    bottom: 3px;
}

.card-suit-top-red {
    font-size: 0.75em;
    position: absolute;
    left: 3px;
    top: -3px;
    color: red;
}

.card-suit-bottom-red {
    font-size: 0.75em;
    position: absolute;
    right: 3px;
    bottom: 3px;
    color: red;
}
</style>