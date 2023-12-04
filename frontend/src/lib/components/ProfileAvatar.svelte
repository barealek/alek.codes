<script>
    let title = "Selvlært Software Udvikler";
    let animating = false;

    setInterval(() => {
        console.log(animating);
    }, 1000);
</script>

<div class="relative flex justify-center items-center">
    <div class="relative w-60 h-60">
        <img class="object-cover w-full h-full rounded-full" src="/assets/me.jpg" alt="alek">
        <span class="absolute bottom-6 right-6 w-5 h-5 rounded-full bg-emerald-500 ring-1 ring-white"></span>
    </div>
</div>

<!-- Name & title -->
<div class="text-center">
    <h1 class="text-4xl font-bold mt-4">Aleksander Rist <span class="text-lg font-normal">(16 år)</span></h1>

    <!-- svelte-ignore a11y-interactive-supports-focus -->
    <span class="text-amber-700/95" on:mouseenter={() => {if (!animating) {animating = true; setTimeout(()=>{animating=false;}, 2000)}}}>
        {#each title.split("") as letter, index (index)}
            {#if letter === " "}
                <span class="inline-block w-1"></span>
            {:else}
                <span style="animation-delay: {50*index}ms;" class="inline-block animate-once {animating ? 'animate-jump' : ''}">{letter}</span>
            {/if}
        {/each}
    </span>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Hedvig+Letters+Sans&display=swap');

    h1 {
        font-family: 'Hedvig Letters Sans', sans-serif;
    }
</style>