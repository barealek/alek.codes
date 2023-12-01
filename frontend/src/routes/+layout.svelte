<script>
    import "../app.css"

    import Navbar from "$lib/components/Navbar.svelte";

    import { onNavigate } from '$app/navigation';

    onNavigate((navigation) => {
        if (!document.startViewTransition) return;

        return new Promise((resolve) => {
            document.startViewTransition(async () => {
                resolve();
                await navigation.complete;
            });
        });
    });

</script>

<Navbar/>

<slot />