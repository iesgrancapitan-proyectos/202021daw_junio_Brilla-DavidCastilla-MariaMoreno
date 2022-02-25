<script>
    import auth from "utils/auth";
    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";

    let token = new URLSearchParams(window.location.search).get("token");
    if (!token) navigate("/");

    let fetching = new Promise(() => {});

    onMount(async () => {
        fetching = await fetch(API_URL + "/user/activate", {
            method: "POST",
            body: JSON.stringify({ token }),
        });
        let { username } = await fetching.json();
        $auth = username;

        navigate("/");
    });
</script>

<section>
    {#await fetching}
        <h2>Activando usuario...</h2>
    {:then _}
        <h2>Redireccionando...</h2>
    {/await}
</section>

<style>
</style>
