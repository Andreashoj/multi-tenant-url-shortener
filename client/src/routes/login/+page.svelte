<script lang="ts">
    import {goto} from "$app/navigation";

    let email = $state('')
    let password = $state('')

    async function submit() {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            }),
        })

        console.log(response.ok)
        if (response.ok) {
            const data = await response.json()
            // Handle success (redirect, etc.)
            // eslint-disable-next-line svelte/no-navigation-without-resolve
            goto("/dashboard")
        } else {
            // Handle error
            console.error('Login failed')
        }
    }

    async function authorizedRequest() {
        const response = await fetch('/api/auth/test', {
            method: 'GET',
            credentials: 'include'
        })

        console.log(await response.json())
    }
</script>

<main class="w-full h-screen flex justify-center items-center flex-col">
    <section>
        <div class="flex flex-col">
            <label for="name">Email</label>
            <input name="email" bind:value={email} type="text">
        </div>
        <div class="flex flex-col mt-4">
            <label for="password">Password</label>
            <input name="password" bind:value={password} type="password">
        </div>

        <button class="px-4 py-2.5 border border-black mt-4" onclick={submit}>
            Login
        </button>

        <button onclick={authorizedRequest}>Make request!</button>
    </section>
</main>
