<script lang="ts">
    import CreateTenantDialog from "$lib/components/dialogs/create-tenant-dialog.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
	import ListItem from "$lib/components/ui/list-item.svelte";
    import {tenantStore} from "../../stores/tenantStore";

    let { data } = $props()
    let { tenants } = $derived($tenantStore)

    tenantStore.getAll()
</script>

<Card.Root>
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

<Card.Root class="mt-6">
    <Card.Header>
        <Card.Title>Tenants</Card.Title>
    </Card.Header>

    <Card.Content>
        <div class="flex flex-col">
            {#each tenants as tenant (tenant.id)}
                <ListItem
                        name={tenant.name}
                        technologies={["Docker", "Svelte", "Go"]}
                        onDelete={() => tenantStore.remove(tenant.id)}
                />
            {/each}
        </div>
    </Card.Content>
</Card.Root>
