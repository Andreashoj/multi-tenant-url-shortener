<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import {Badge} from "$lib/components/ui/badge";
    import {EllipsisVertical} from "@lucide/svelte";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";

    let { onDelete, name, technologies }: { onDelete: () => void,  name: string, technologies: string[] } = $props()

</script>


<Card.Root class="mt-6">
    <Card.Content class="flex justify-between">
        <div>
            <h3 class="first-letter:capitalize">
                { name }
            </h3>
            <div class="flex gap-2 mt-3">
                {#each technologies as tech (tech)}
                    <Badge variant="secondary" class="bg-gray-200">
                        { tech }
                    </Badge>
                {/each}
            </div>
        </div>


        <DropdownMenu.Root>
            <DropdownMenu.Trigger>
                {#snippet child({props})}
                    <EllipsisVertical { ...props } />
                {/snippet}
            </DropdownMenu.Trigger>
            <DropdownMenu.Content class="w-56" align="start">
                <DropdownMenu.Label>Tenant</DropdownMenu.Label>
                <DropdownMenu.Group>
                    <DropdownMenu.Item onclick={onDelete}>
                        Delete
                        <DropdownMenu.Shortcut>⇧⌘D</DropdownMenu.Shortcut>
                    </DropdownMenu.Item>
                </DropdownMenu.Group>
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    </Card.Content>
</Card.Root>
