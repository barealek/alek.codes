<script>
    export let status = 'offline';
</script>

<div class="relative flex justify-center items-center">
    <div class="relative w-60 h-60">
        <img class="object-cover w-full h-full rounded-full" src="/assets/me.jpg" alt="alek">

        <span class="absolute bottom-6 right-6 w-5 h-5 rounded-full {status !== 'online' ? 'bg-neutral-400' : 'bg-emerald-500'} ring-1 ring-white hover:p-visible">
            <p class="opacity-0 invisible duration-300 absolute flex items-center justify-center w-40 p-3 text-gray-600 bg-white rounded-lg shadow-lg -left-[18px] -top-16 shadow-gray-200">
                <span class="truncate ">Jeg er {status === 'online' ? 'online 😄' : 'offline 😢'}</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 absolute rotate-45 left-4 bottom-0.5 -mb-3 transform text-white dark:text-gray-800 fill-current" stroke="currentColor" viewBox="0 0 24 24" >
                    <path d="M20 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"></path>
                </svg>
            </p>
        </span>

    </div>
</div>

<!-- Name & title -->
<div class="text-center">
    <h1 class="text-4xl font-bold mt-4">Aleksander Rist <span class="text-lg font-normal">(16 år)</span></h1>
    <p class="text-amber-600/80">Selvlært Software Udvikler</p>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Hedvig+Letters+Sans&display=swap');

    h1 {
        font-family: 'Hedvig Letters Sans', sans-serif;
    }

    .hover\:p-visible:hover > p {
        opacity: 1;
        visibility: visible;
    }
</style>