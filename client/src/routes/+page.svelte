<script lang="ts">
    let email = $state('')
    let password = $state('')

    async function submit() {
        const response = await fetch('http://localhost:8080/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            }),
            credentials: 'include' // Important for cookies!
        })

        if (response.ok) {
            const data = await response.json()
            // Handle success (redirect, etc.)
            console.log(data)
        } else {
            // Handle error
            console.error('Login failed')
        }
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
    </section>
</main>
