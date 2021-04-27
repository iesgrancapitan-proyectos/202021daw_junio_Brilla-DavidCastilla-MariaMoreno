<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";

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
        />
    {/each}
</main>

<button on:click={logout}>Logout</button>

<section>
    <textarea name="brillo" id="brillo" cols="30" rows="10" onkeyup="countChars(this);" />
    <div>
        <span>0/250</span>
        <input type="submit" value="crear brillo" />
    </div>
</section>

<style lang="scss">
    main {
        display: grid;
        grid-gap: 16px;
        margin: 32px;
    }

    section {
        margin: 32px;
        text-align: center;
        textarea {
            border: solid 2px #f9c55f;
            resize: none;
        }

        div {
            display: flex;
            justify-content: space-between;
            margin-top: 16px;
        }
    }
</style>
