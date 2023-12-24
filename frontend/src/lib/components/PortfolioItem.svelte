<script>
    export let portfolio;
    let title = portfolio.title;
    let type = portfolio.type || "???";
    let description = portfolio.description;
    let image = portfolio.image;
    let sublink = portfolio.sublink
    
</script>

<div class="max-w-xl min-w-lg overflow-hidden bg-white rounded-lg shadow-md dark:bg-gray-800">
    <img class="aspect-video object-cover w-full h-64" src={image} alt="Billed af {type}">

    <div class="p-6">
        <div>
            <span class="text-xs font-medium text-blue-600 uppercase dark:text-blue-400">{type}</span>
            <a href="#" class="block mt-2 text-xl font-semibold text-gray-800 transition-colors duration-300 transform dark:text-white hover:text-gray-600 hover:underline" tabindex="0" role="link">{title}</a>
            <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">{description}</p>
        </div>

        <div class="mt-4">
            <div class="flex items-center">
                <div class="flex items-center">
                    <img class="object-cover h-10 rounded-full" src="https://images.unsplash.com/photo-1586287011575-a23134f797f9?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=48&q=60" alt="Avatar">
                    <a href="#" class="mx-2 font-semibold text-gray-700 dark:text-gray-200" tabindex="0" role="link">Jone Doe</a>
                </div>
                <span class="mx-1 text-xs text-gray-600 dark:text-gray-300">21 SEP 2015</span>
            </div>
        </div>
    </div>
</div>