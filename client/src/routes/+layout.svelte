<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
    import {goto} from "$app/navigation";

	let { children, data } = $props();

</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>


{@render children?.()}
