<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";
    import ArrowLeft from "svelte-material-icons/ArrowLeft";

    export let id;

    let bright;
    let comments;

    onMount(async () => {
        let data = await fetch(API_URL + `/brights/${id}`);
        bright = await data.json();
    });
</script>

<section>
    <button on:click={() => history.back()}> <ArrowLeft /></button>

    {#if bright}
        <Brillo
            user={{
                username: bright.username,
                name: bright.name,
                img: `/media/${bright.userKey}/pp`,
            }}
            content={bright.content}
            uploadDate={new Date(bright.created_at)}
            id={bright._key}
            rebrillos={bright.rebrillos}
            interactions={bright.interactions}
            ncomments={bright.comments.length}
            media={bright.media}
        />
    {/if}
    {#if bright && bright.comments.length > 0}
        {#each bright.comments as comment}
            <Brillo
                user={{
                    username: comment.username,
                    name: comment.name,
                    img: `/media/${comment.userKey}/pp`,
                }}
                content={comment.content}
                uploadDate={new Date(comment.created_at)}
                id={comment._key}
                rebrillos={comment.rebrillos}
                interactions={comment.interactions}
                ncomments={comment.comments}
                media={comment.media}
            />
        {/each}
    {/if}
</section>

<style lang="scss">
    section {
        margin: 32px;
        button {
            background: none;
            border: 0;
            font-size: 1.5rem;
            cursor: pointer;
        }
    }
</style>
