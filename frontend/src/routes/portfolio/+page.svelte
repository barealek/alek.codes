<script>
    
    import PortfolioItem from '$lib/components/PortfolioItem.svelte';

    import portfolios from '$lib/data/portfolio.js';
    
</script>

<div class="h-screen w-screen flex justify-center items-center md:mt-0 mt-64">
    <div>
        <section class="bg-white dark:bg-gray-900 mt-12">
            <div class="container px-6 py-10 mx-auto">
                <h1 class="text-2xl font-semibold text-center text-gray-800 capitalize lg:text-3xl dark:text-white">Portfolio
                </h1>
        
                <p class="mt-4 text-center text-gray-500 dark:text-gray-300">
                    Nedenunder er en liste af løsninger jeg har udviklet. Klik på en af dem for at læse mere.
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