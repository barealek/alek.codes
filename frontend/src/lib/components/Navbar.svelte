<script>
    import { page } from '$app/stores';

    let path = $page.url.pathname
    
    let pages = {
        hjem: "/",
        // blogs: "/blogs/",
        portfolie: "/portfolio/"
    }

    let activeClasses = "border-b-2 border-amber-500"
    let unactiveClasses = "border-b-2 border-transparent hover:border-amber-500"

</script>

<nav class="w-screen bg-white shadow-sm transition-all hover:shadow dark:bg-gray-800 flex h-16 md:-mb-16">
    <div class="container flex items-center justify-center p-6 md:gap-0 gap-4 mx-auto text-gray-600">

        {#each Object.entries(pages) as [key, value]}
            <a on:click={() => {path = value}} href={value} class="{path === value ? activeClasses : unactiveClasses} hover:text-gray-800 transition-colors duration-300 transform dark:hover:text-gray-200 mx-1.5 sm:mx-6">{key}</a>
        {/each}
    </div>
</nav>