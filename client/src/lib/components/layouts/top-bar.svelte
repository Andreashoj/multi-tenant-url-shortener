<script lang="ts">
    import Button from "../ui/button/button.svelte";
    import {logout} from "../../../api/auth";
    import {goto} from "$app/navigation";

    let { title } = $props()

    async function handleLogout() {
        try {
            await logout()
            await goto("/login")
        } catch(e) {
            console.error(`Logout failed: ${e}`)
        }
    }
</script>

<header class="border-b-2 border-b-slate-500 capitalize min-h-16 p-4 flex justify-between">
    <h1 class="text-xl font-semibold leading-10 text-slate-800">
        { title }
    </h1>

    <Button class="cursor-pointer bg-gray-200 text-black" onclick={handleLogout}>
        Logout (should be moved to profile icon)
    </Button>
</header>
