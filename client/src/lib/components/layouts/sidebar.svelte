<script lang="ts">
    import {page} from "$app/state";
    import {CircleUser, LayoutDashboard, Settings} from "@lucide/svelte";
    function isActive(path: string) {
        return page.url.pathname === path;
    }
</script>
<div class="w-20 h-screen"></div>
<nav class="bg-emerald-950 w-20 pt-4 pb-8 flex flex-col items-center fixed h-screen">
    <span class="text-white">Logo</span>

    <div class="flex flex-col gap-2 items-center mt-12">
        <!-- TURN TO COMPONENT -->
        <a
            href="/dashboard"
            class="text-white rounded p-2 {isActive('/dashboard') ? 'bg-emerald-900' : '' }"
        >
            <LayoutDashboard />
        </a>
        <a
            href="/settings"
            class="text-white rounded p-2  {isActive('/settings') ? 'bg-emerald-900' : '' }"
        >
            <Settings />
        </a>
    </div>

    <button class="mt-auto text-white">
        <CircleUser />
    </button>
</nav>
