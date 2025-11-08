<script lang="ts">
    import {goto} from "$app/navigation";
    import {login} from "../../api/auth";
    import {Button} from "$lib/components/ui/button";
    import {Label} from "$lib/components/ui/label";
    import {Input} from "$lib/components/ui/input";
    import * as Card from "$lib/components/ui/card/index"

    let email = $state('')
    let password = $state('')

    async function submit() {
        try {
            await login(email, password)
            await goto("/dashboard")
        } catch (e) {
            console.error(`Login failed: ${e}`)
        }
    }
</script>

<Card.Root class="w-full max-w-sm">
    <Card.Header>
        <Card.Title>Login to your account</Card.Title>
        <Card.Description
        >Enter your email below to login to your account</Card.Description
        >
    </Card.Header>
    <Card.Content>
        <div class="flex flex-col">
            <Label class="mb-2" for="name">Email</Label>
            <Input name="email" bind:value={email} type="text" />
        </div>
        <div class="flex flex-col mt-4">
            <Label class="mb-2" for="password">Password</Label>
            <Input name="password" bind:value={password} type="password" />
        </div>

        <Card.Action>
            <Button class="px-4 py-2.5 border border-black mt-4" onclick={submit}>
                Login
            </Button>
        </Card.Action>
    </Card.Content>
</Card.Root>