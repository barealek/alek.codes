<script>
    let activeClasses = "border-b-2 border-blue-500"
    let unactiveClasses = "border-b-2 border-transparent hover:border-blue-500"

    let path = window.location.pathname
    let activePage = path
    
    let pages = {
        home: "/",
        blogs: "/blogs/",
        contact: "/contact/"
    }

</script>
<nav class="absolute w-screen bg-white shadow-sm transition-all hover:shadow dark:bg-gray-800 flex mb-8">
    <div class="container flex items-center justify-center p-6 mx-auto text-gray-600">

        {#each Object.entries(pages) as [key, value]}
            <a on:click={() => {activePage = value}} href={value} class="{activePage === value ? activeClasses : unactiveClasses} hover:text-gray-800 transition-colors duration-300 transform dark:hover:text-gray-200 mx-1.5 sm:mx-6">{key}</a>
        {/each}
    </div>
</nav>