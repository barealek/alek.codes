<script>

    export let portfolio

    let title = portfolio.title
    let type = portfolio.type || "???"
    let description = portfolio.description
    let image = portfolio.image
    let sublink = portfolio.sublink
    let date = portfolio.date || "???"
    
</script>

<div class="max-w-xl min-w-lg overflow-hidden bg-white rounded-lg shadow-md">
    <img class="aspect-[2/1] object-cover w-full h-48" src={image} alt="Billed af {type}">

    <div class="p-6">
        <div>
            <span class="text-xs font-medium text-amber-600 uppercase">{type}</span>
            <a href={sublink} class="block mt-2 text-xl font-semibold text-gray-800 transition-colors duration-300 transform hover:text-gray-600 hover:underline" tabindex="0">{title}</a>
            <p class="mt-2 text-sm text-gray-600 ">{description}</p>
        </div>

        <div class="mt-4">
            <div class="flex items-center">
                <div class="flex items-center text-gray-600">
                    <svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                </div>
                <span class="mx-1 text-xs text-gray-600 ">{date}</span>
            </div>
        </div>
    </div>
</div>