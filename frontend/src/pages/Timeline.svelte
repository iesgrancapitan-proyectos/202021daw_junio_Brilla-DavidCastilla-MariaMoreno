<script>
    import { onMount } from "svelte";

    import auth from "utils/auth";

    let brights = [];

    async function logout() {
        try {
            await fetch(API_URL + "/logout", { credentials: "include" });
            $auth = null;
        } catch (e) {}
    }

    onMount(async () => {
        let data = await fetch(API_URL + "/timeline");
        brights = [...brights, ...(await data.json())];
    });
</script>

<h1>Timeline</h1>

{#each brights as bright}
    <article>
        <h1>@{bright.username}</h1>
        <p>{bright.content}</p>
        <span>{new Date(bright.created_at).toLocaleString("es")}</span>
    </article>
{/each}

<button on:click={logout}>Logout</button>
