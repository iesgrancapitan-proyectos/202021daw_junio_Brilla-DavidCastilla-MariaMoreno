<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";

    export let id;

    let bright;

    onMount(async () => {
        let data = await fetch(API_URL + `/brights/${id}`);
        bright = await data.json();
    });
</script>

<section>
    {#if bright}
        <Brillo
            user={{
                username: bright.username,
                name: bright.name,
                profile_img: bright.profile_img,
            }}
            content={bright.content}
            uploadDate={new Date(bright.created_at)}
            id={bright._key}
            rebrillos={bright.rebrillos}
            interactions={bright.interactions}
            comments={bright.comments}
            media={bright.media}
        />
    {/if}
</section>

<style>
    section {
    }
</style>
