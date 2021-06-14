<script>
    import { onMount } from "svelte";
    import auth from "utils/auth";
    import Brillo from "../components/Brillo.svelte";
    import Settings from "svelte-material-icons/Settings";
    import ContentSave from "svelte-material-icons/ContentSave";
    import Menu from "../components/Menu.svelte";
    import InfiniteLoading from "svelte-infinite-loading";

    let name = "";
    export let username = "";
    let bio = "",
        followed = "",
        followers = "",
        nbrillos = "",
        imgPerfil = "",
        key = "",
        newUsername = username,
        files = {};

    let brights = [];
    let isFollowing = false;
    let edits = true;

    onMount(async () => {
        let res = await fetch(API_URL + `/user/${username}`);
        let info = await res.json();

        key = info["_key"];
        bio = info.bio;
        name = info.name;
        imgPerfil = info.imgPerfil;
        imgPerfil = `/media/${key}/pp`;
        //n brillos suma de brillos.
        nbrillos = await fetch(API_URL + `/user/${key}/brights/count`);
        let nbrillosJson = await nbrillos.json();
        nbrillos = nbrillosJson["nbrillos"];

        followed = await fetch(API_URL + `/nfollowed/${key}`);
        let followedJson = await followed.json();
        followed = followedJson["followed"];

        followers = await fetch(API_URL + `/nfollowers/${key}`);
        let followersJson = await followers.json();
        followers = followersJson["followers"];

        let following = await fetch(API_URL + `/user/${key}/isfollowing`);
        let { follow } = await following.json();
        isFollowing = follow;

        fetchBrights(0);
    });

    async function follow() {
        let res_follow = await fetch(API_URL + `/user/${key}/follow`, {
            method: "PUT",
            credentials: "include",
        });
        let json = await res_follow.json();

        isFollowing = json.result;
    }

    async function edit() {
        console.dir(files);
        //edit
        let form = new FormData();
        form.append("bio", bio);
        form.append("username", newUsername);
        form.append("name", name);

        if (files[0]) form.append("profile", files[0]);

        let res_edit = await fetch(API_URL + `/user/edit`, {
            method: "POST",
            body: form,
            credentials: "include",
        });

        if (res_edit.status == 200) {
            history.replaceState("", "", `/user/${newUsername}`);
            $auth = newUsername;
        } else if (res_edit.status == 409) {
            newUsername = username;
        }

        edits = !edits;
    }

    async function fetchBrights(offset, event) {
        let data = await fetch(
            API_URL + `/user/${key}/brights?offset=${offset}`
        );

        if (data.status == 404) {
            return;
        }

        let jsonData = await data.json();
        // jsonData.forEach((el) => brights.push(el));
        // brights = brights;
        brights = [...brights, ...jsonData];

        console.log(brights);
        if (event != null && jsonData.length < 10) event.detail.complete();
        event?.detail.loaded();
    }

    function changePicture() {
        let pic = new FileReader();
        pic.onload = () => (imgPerfil = pic.result);
        pic.readAsDataURL(files[0]);
    }
</script>

<main>
    <Menu />
    <div class="inicio">
        <!-- <svg height="60%" width="100%">
            <circle
                cx="50%"
                cy="-90%"
                r="85%"
                stroke="#f9c55f"
                stroke-width="3"
                fill="#f9c55f"
            />
        </svg> -->

        <div>
            <!-- <p>@{username}</p> -->
            <p>
                @<input
                    type="text"
                    bind:value={newUsername}
                    disabled={edits}
                    size={newUsername.length}
                />
            </p>

            <img src={imgPerfil} alt="img perfil" />
            {#if !edits}<input
                    type="file"
                    bind:files
                    on:change={changePicture}
                />
            {/if}

            <!-- <p>{name}</p> -->
            <input type="text" bind:value={name} disabled={edits} />
        </div>
    </div>

    <div />

    <div class="info">
        <textarea type="text" bind:value={bio} disabled={edits} />

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
</main>

<section>
    {#each brights as bright}
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
            ncomments={bright.comments}
            media={bright.media}
            rebrightedBy={username}
        />
    {/each}

    <InfiniteLoading
        distance={200}
        on:infinite={(e) => fetchBrights(brights.length, e)}
    >
        <div slot="spinner" />
    </InfiniteLoading>
</section>

<style lang="scss">
    @svg circle {
        @circle {
            cx: 50%;
            cy: var(--center-y);
            r: var(--radius);
            fill: #f9c55f;
        }
    }

    main {
        // padding: 16px;
        > :global(button) {
            position: absolute;
            left: 0 from-md(10%) from-lg(20%);
            background: none;
            border: 0;
            padding: 16px;
            font-size: 1.5em;
        }
        .inicio {
            text-align: center;
            width: 100%;
            display: flex;
            justify-content: center;
            position: relative;
            background: svg(circle param(--center-y -75%) param(--radius 60%))
                center / cover
                from-md(
                    svg(circle param(--center-y -150%) param(--radius 40%))
                        center / cover
                )
                from-lg(
                    svg(circle param(--center-y -200%) param(--radius 35%))
                        center / cover
                );

            div {
                display: flex;
                flex-direction: column;
                align-items: center;
                z-index: 0;
                top: 8vh;
                margin: 0 auto;
                width: 50%;

                img {
                    border-radius: 16px;
                    size: 24vmin from-md(12vmin);
                    object-fit: cover;
                }

                p {
                    display: flex;
                    justify-content: center;
                }
            }
        }
        .info {
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            justify-content: space-around;
            width: from-md(80%) from-lg(60%);
            margin: 0 auto;
            padding: 16px;
            input[type="text"],
            textarea {
                flex-basis: 68%;
                outline: none;
                text-align: start;
            }

            div {
                flex-basis: 30%;
                text-align: end;
            }
            hr {
                flex-basis: 100%;
            }
            button {
                background: var(--primary-color);
                border: 0;
                padding: 8px 16px;
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

        *:disabled {
            all: unset;
            // width: 100%;
            text-align: center;
        }
    }

    section {
        padding: 16px;
        width: from-md(80%) from-lg(60%);
        margin: 0 auto;
    }
</style>
