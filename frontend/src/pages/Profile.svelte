<script>
    import { onMount } from "svelte";
    import auth from "utils/auth";
    import Brillo from "../components/Brillo.svelte";
    import Settings from "svelte-material-icons/Settings";
    import ContentSave from "svelte-material-icons/ContentSave";
    import Menu from "../components/Menu.svelte";

    let name = "";
    export let username = "";
    let bio = "",
        followed = "",
        followers = "",
        nbrillos = "",
        imgPerfil = "";

    let brights = [];
    let isFollowing = false;
    let edits = true;

    onMount(async () => {
        let res = await fetch(API_URL + `/user/${username}`);
        let info = await res.json();

        bio = info.bio;
        name = info.name;
        imgPerfil = info.imgPerfil;
        imgPerfil = "https://picsum.photos/200";
        //n brillos suma de brillos.
        nbrillos = await fetch(API_URL + `/user/${username}/brights/count`);
        let nbrillosJson = await nbrillos.json();
        nbrillos = nbrillosJson["nbrillos"];

        followed = await fetch(API_URL + `/nfollowed/${username}`);
        let followedJson = await followed.json();
        followed = followedJson["followed"];

        followers = await fetch(API_URL + `/nfollowers/${username}`);
        let followersJson = await followers.json();
        followers = followersJson["followers"];

        let data = await fetch(API_URL + `/user/${username}/brights`);
        brights = [...brights, ...(await data.json())];

        let following = await fetch(API_URL + `/user/${username}/isfollowing`);
        let { follow } = await following.json();
        isFollowing = follow;
    });

    async function follow() {
        let res_follow = await fetch(API_URL + `/user/${username}/follow`, {
            method: "PUT",
            credentials: "include",
        });
        let json = await res_follow.json();

        isFollowing = json.result;
    }

    async function edit() {
        //edit
        let form = new FormData();
        form.append("bio", bio);
        form.append("username", username);
        form.append("name", name);
        // form.append("profile_img", profile_img);

        let res_edit = await fetch(API_URL + `/user/edit`, {
            method: "POST",
            body: form,
            credentials: "include",
        });

        edits = !edits;
    }
</script>

<section>
    <Menu />

    <div class="inicio">
        <svg height="60vw" width="100vw">
            <circle
                cx="50vw"
                cy="-25vw"
                r="60vw"
                stroke="#f9c55f"
                stroke-width="3"
                fill="#f9c55f"
            />
        </svg>
        <div>
            <!-- <p>@{username}</p> -->
            <p>
                @<input type="text" bind:value={username} disabled={edits} />
            </p>

            <img src={imgPerfil} alt="img perfil" />

            <!-- <p>{name}</p> -->
            <input type="text" bind:value={name} disabled={edits} />
        </div>
    </div>

    <div />

    <div class="info">
        <input type="text" bind:value={bio} disabled={edits} />

        <div>
            <p>{followed} Followed</p>
            <p>{followers} followers</p>
            <p>{nbrillos} Brillos</p>
        </div>
        {#if username != $auth}
            {#if isFollowing}
                <button on:click={follow}>Following</button>
            {:else}
                <button on:click={follow}>Follow</button>
            {/if}
        {:else if edits}
            <button on:click={() => (edits = !edits)}>
                Edit <Settings />
            </button>
        {:else}
            <button on:click={edit}>
                Save <ContentSave />
            </button>
        {/if}

        <hr />
    </div>
</section>

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

<style lang="scss">
    section {
        margin: 16px;
        .inicio {
            text-align: center;
            width: 100%;
            display: flex;
            justify-content: center;

            svg {
                position: relative;
                z-index: -1;
                top: 0px;
                left: 0px;
            }

            div {
                position: absolute;
                z-index: 0;
                top: 8vh;
                margin: 0 auto;
                width: 50%;

                img {
                    border-radius: 15%;
                }

                p {
                    display: flex;
                }
            }
        }
        .info {
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            justify-content: space-around;
            input[type="text"] {
                flex-basis: 68%;
            }

            div {
                flex-basis: 30%;
            }
            hr {
                flex-basis: 100%;
            }
            button {
                background: var(--primary-color);
                border: 1px solid var(--dark-primary-color);
                padding: 8px;
                // box-shadow: 2px 2px 2px 2px var(--dark-primary-color);
                margin-bottom: 16px;
                border-radius: 12px;
                display: flex;
                align-items: center;
                cursor: pointer;
                :global(svg) {
                    padding: 1px;
                    margin: 3px;
                }
            }
        }

        input:disabled {
            all: unset;
            width: 100%;
            text-align: center;
        }
    }
</style>
