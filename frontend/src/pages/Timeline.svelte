<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";
    import FormBrillo from "components/FormCreateBrillo";

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

<main>
    {#each brights as bright}
        <Brillo
            user={{
                username: bright.username,
                name: bright.name,
                profile_img: bright.profile_img,
            }}
            content={bright.content}
            uploadDate={new Date(bright.created_at)}
            id={bright.id}
        />
    {/each}

    <FormBrillo />
</main>

<button on:click={logout}>Logout</button>

<style lang="scss">
    main {
        display: grid;
        grid-gap: 16px;
        margin: 32px;
    }
</style>
