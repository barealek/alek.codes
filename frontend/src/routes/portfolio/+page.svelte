<script>
    
    import PortfolioItem from '$lib/components/PortfolioItem.svelte';
</script>

<div class="h-screen w-screen flex justify-center items-center">
    <div>
        <section class="bg-white dark:bg-gray-900">
            <div class="container px-6 py-10 mx-auto">
                <h1 class="text-2xl font-semibold text-center text-gray-800 capitalize lg:text-3xl dark:text-white">Portfolio
                </h1>
        
                <p class="mt-4 text-center text-gray-500 dark:text-gray-300">
                    TODO: Tilføj portfolie ;)
                </p>
        
        
                <div class="grid grid-cols-1 gap-8 mt-8 xl:mt-12 xl:gap-12 md:grid-cols-2 xl:grid-cols-3">
                    {#each portfolios as poitem}
                        <PortfolioItem portfolio={poitem}/>
                    {/each}
                </div>
            </div>
        </section>
    </div>
</div>