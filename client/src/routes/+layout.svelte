<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
    import Sidebar from "$lib/components/layouts/sidebar.svelte";
    import TopBar from "$lib/components/layouts/top-bar.svelte";
    import { page } from "$app/state"

	let { children, data } = $props();
    let pageTitle = $derived(() => page.route.id?.replace("/", "") || "Dashboard")
    let showDashboardUI = $derived(() => data.user)
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>


{#if showDashboardUI()}
    <div class="flex w-full min-h-screen">
        <Sidebar />
        <main class="flex flex-col flex-1 bg-gray-50">
            <TopBar title={pageTitle()} />
            <div class="px-8 py-12">
                {@render children?.()}
            </div>
        </main>
    </div>
{:else}
    <main class="w-full h-screen flex justify-center items-center flex-col">
        {@render children?.()}
    </main>
{/if}
