<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";

    export let id;

    onMount(async () => {
        let data = await fetch(API_URL + `/brillo/${id}`);
        brights = [...brights, ...(await data.json())];
        console.log(brights);
    });
</script>

<section>
    {#each brights as bright}
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
    {/each}
</section>

<style>
    section {
    }
</style>
