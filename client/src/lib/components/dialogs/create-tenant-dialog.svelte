<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import * as Dialog from "$lib/components/ui/dialog/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import {tenantStore} from "../../../stores/tenantStore";

    let name = $state("")

</script>

<Dialog.Root>
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
                {name}
            </div>
        </div>
        <Dialog.Footer>
            <Button type="submit" onclick={() => tenantStore.create(name)}>Save changes</Button>
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>