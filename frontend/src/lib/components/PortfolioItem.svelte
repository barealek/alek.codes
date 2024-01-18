<script>

    export let portfolio

    let title = portfolio.title
    let type = portfolio.type || "???"
    let description = portfolio.description
    let links = portfolio.links
    let image = portfolio.image
    let date = portfolio.date || "???"

    let icons = {
        github: `<svg class="w-5 h-5 sm:w-6 sm:h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path></svg>`,
        hjemmeside: `<svg class="w-5 h-5 sm:w-6 sm:h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>`
    }
    
</script>

<div class="grid max-w-xl min-w-lg overflow-hidden bg-white rounded-lg shadow-md">
    <img class="aspect-[2/1] object-cover w-full h-48" src={image} alt="Billed af {type}">

    <div class="p-6">
        <div>
            <span class="text-xs font-medium text-amber-600 uppercase">{type}</span>
            <span class="block mt-2 text-xl font-semibold text-gray-800 transition-colors duration-300 transform hover:text-gray-600 hover:underline">{title}</span>
            <p class="min-h-[5rem] mt-2 text-sm text-gray-600 ">{description}</p>
        </div>

        <div class="mt-4">
            <div class="flex items-center">
                <div class="flex items-center text-gray-600">
                    <svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                </div>
                <span class="mx-1 text-xs text-gray-600 ">{date}</span>
            </div>
        </div>

        <div class=" self-end mt-2">
            <div class="flex items-center justify-evenly text-gray-600">
                {#each Object.keys(links) as key}
                    {#if links[key]}
                        <a href={links[key]} target="_blank" rel="noopener noreferrer">
                            {@html icons[key] || key}
                        </a>
                    {/if}
                {/each}
            </div>
        </div>
    </div>
</div>