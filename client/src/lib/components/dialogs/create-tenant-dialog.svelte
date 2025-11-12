<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import * as Dialog from "$lib/components/ui/dialog/index.js";
    import * as Select from "$lib/components/ui/select/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import {tenantStore} from "../../../stores/tenantStore";

    let name = $state("")
    let dialogOpen = $state(false)
    const dbOptions = [
        { value: "shared", label: "Shared" },
        { value: "isolated", label: "Isolated" },
        { value: "schema", label: "Schema" },
    ];

    let DBType = $state("");

    function create() {
        tenantStore.create(name, DBType)
        dialogOpen = false
        name = ""
    }

    const triggerContent = $derived(
        dbOptions.find((f) => f.value === DBType)?.label ?? "Select a database type"
    );
</script>

<Dialog.Root bind:open={dialogOpen}>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Create tenant
    </Dialog.Trigger>
    <Dialog.Content class="sm:max-w-[425px]">
        <Dialog.Header>
            <Dialog.Title>Tenant</Dialog.Title>
        </Dialog.Header>
        <div class="grid gap-4 py-4">
            <div class="grid grid-cols-4 items-center gap-4">
                <Label for="name" class="text-right">Name</Label>
                <Input id="name" bind:value={name} class="col-span-3" />
            </div>

            <div class="flex justify-between">
                <Label for="databaseSelection" class="text-right">Database</Label>
                <Select.Root type="single" name="databaseSelection" bind:value={DBType}>
                    <Select.Trigger>
                        {triggerContent}
                    </Select.Trigger>
                    <Select.Content>
                        <Select.Group>
                            <Select.Label>Database Type</Select.Label>
                            {#each dbOptions as db (db.value)}
                                <Select.Item
                                        value={db.value}
                                        label={db.label}
                                >
                                    {db.label}
                                </Select.Item>
                            {/each}
                        </Select.Group>
                    </Select.Content>
                </Select.Root>
            </div>
        </div>
        <Dialog.Footer>
            <Button type="submit" onclick={create}>Create</Button>
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>