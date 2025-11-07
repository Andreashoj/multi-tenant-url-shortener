<script lang="ts">
    import {goto} from "$app/navigation";
	import Button from "$lib/components/ui/button/button.svelte";
    import CreateTenantDialog from "$lib/components/dialogs/create-tenant-dialog.svelte";
    import * as Card from "$lib/components/ui/card/index.js";

    let { data } = $props()

    async function logout() {
        const res = await fetch("/api/auth/logout", {
            method: "POST",
            credentials: "include"
        })

        if (res.status == 200) {
            await goto("/login")
        }
    }
</script>

<div class="w-full h-screen relative bg-black p-8">
    <Button class="absolute right-8 top-5 cursor-pointer" variant="secondary" onclick={logout}>
        Logout
    </Button>

    <Card.Root class="mt-12">
        <Card.Header>
            <Card.Title>User: {data.user?.email}</Card.Title>
            <Card.Title>Role: {data.user?.role}</Card.Title>
        </Card.Header>
        <Card.Content>
            <div>
                <CreateTenantDialog />
            </div>
        </Card.Content>
    </Card.Root>

    <Card.Root class="mt-12">
        <Card.Header>
            <Card.Title>Tenants</Card.Title>
        </Card.Header>

        <Card.Content>
            <div>
                <Button>Create Tenant</Button>
            </div>
        </Card.Content>
    </Card.Root>
</div>
