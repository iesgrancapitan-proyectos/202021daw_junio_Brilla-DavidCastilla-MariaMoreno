<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";

    export let id;

    let bright;
    let comments;

    onMount(async () => {
        let data = await fetch(API_URL + `/brights/${id}`);
        bright = await data.json();
        console.log(bright.comments);
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
            comments={bright.comments.length}
            media={bright.media}
        />
    {/if}
    {#if bright && bright.comments.length > 0}
        {#each bright.comments as comment}
            <Brillo
                user={{
                    username: comment.username,
                    name: comment.name,
                    profile_img: comment.profile_img,
                }}
                content={comment.content}
                uploadDate={new Date(comment.created_at)}
                id={comment._key}
                rebrillos={comment.rebrillos}
                interactions={comment.interactions}
                comments={comment.comments}
                media={comment.media}
            />
        {/each}
    {/if}
</section>

<style>
    section {
        margin: 32px;
    }
</style>
