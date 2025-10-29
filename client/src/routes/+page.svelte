<script lang="ts">
    let email = $state('')
    let password = $state('')
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

        <button class="px-4 py-2.5 border border-black mt-4">
            Login
        </button>
    </section>
</main>
