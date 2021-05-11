<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";
    import FormBrillo from "components/FormCreateBrillo";
    import auth from "utils/auth";
    import PencilPlus from "svelte-material-icons/PencilPlus";
    import Close from "svelte-material-icons/Close";

    let brights = [];
    let see = false;
    async function logout() {
        try {
            await fetch(API_URL + "/logout", { credentials: "include" });
            $auth = null;
        } catch (e) {}
    }

    onMount(async () => {
        let data = await fetch(API_URL + "/timeline");
        brights = [...brights, ...(await data.json())];
        console.log(brights);
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
            id={bright._key}
            rebrillos={bright.rebrillos}
            interactions={bright.interactions}
            comments={bright.comments}
            media={bright.media}
        />
    {/each}

    <div class:active={see}>
        <button on:click={() => (see = false)}><Close /></button>
        <FormBrillo route="/brights" />
    </div>

    <button on:click={() => (see = true)}><PencilPlus /></button>
</main>

<button on:click={logout}>Logout</button>

<style lang="scss">
    main {
        display: grid;
        grid-gap: 16px;
        margin: 32px;
        > button {
            position: fixed;
            bottom: 16px;
            right: 16px;
            background: var(--primary-color);
            border: 0;
            padding: 16px;
            border-radius: 100vmax;
            :global(svg) {
                display: block;
            }
        }

        div {
            position: fixed;
            bottom: 0;
            transform: translateY(100%);
            text-align: end;

            &.active {
                transform: translateY(0);
                width: 100%;
                right: 0;
                background-color: var(--background-color);
                z-index: 999;
            }
            button {
                border: 0;
                background: var(--background-color);
                font-size: 16px;
                margin: 16px 24px 0 0;
                :global(svg) {
                    color: var(--dark-primary-color);
                    font-size: 24px;
                }
            }
        }
    }
</style>
