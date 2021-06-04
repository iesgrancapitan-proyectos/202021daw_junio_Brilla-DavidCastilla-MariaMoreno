<script>
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";
    import Popover from "svelte-popover";
    import debounce from "lodash/debounce";
    import Brillo from "components/Brillo";
    import Menu from "components/Menu";
    import FormBrillo from "components/FormCreateBrillo";
    import InfiniteLoading from "svelte-infinite-loading";
    import auth from "utils/auth";
    import { truncate, merge } from "utils/strings";
    import PencilPlus from "svelte-material-icons/PencilPlus";
    import Close from "svelte-material-icons/Close";
    import Tumbleweed from "assets/tumbleweed.png";

    let brights = [];
    let userSearch = [];
    let brightSearch = [];
    let see = false;
    async function logout() {
        try {
            await fetch(API_URL + "/logout", { credentials: "include" });
            $auth = null;
        } catch (e) {}
    }

    async function fetchBrights(offset, event) {
        let data = await fetch(API_URL + `/timeline?offset=${offset}`);
        let jsonData = await data.json();
        jsonData.forEach((el) => brights.push(el));
        brights = brights;
        if (event != null && jsonData.length < 10) event.detail.complete();
        event?.detail.loaded();
    }

    const handleInput = debounce(async ({ target: { value: value } }) => {
        userSearch = [];
        let res = await fetch(
            API_URL + `/search?q=${encodeURIComponent(value)}`
        );
        let json = await res.json();
        json.forEach((el) => {
            let t = document.createElement("div");
            t.innerHTML = el.content;
            el.content = t.innerText;
        });
        console.log(json);
        if (value.startsWith("@")) userSearch = json;
        else brightSearch = json;
    }, 300);

    onMount(async () => {
        fetchBrights(0);
        // console.log(brights.length);
    });
</script>

<div class="principal">
    <Menu />
    <div>
        <input on:input={handleInput} />
        <div>
            <ul>
                {#each userSearch as user}
                    <li>
                        <Link to={`/user/${user.username}`}>
                            <div>
                                <h3>@{user.username}</h3>
                            </div>
                        </Link>
                    </li>
                {/each}
                {#each brightSearch as bright}
                    <li>
                        <Link to={`/brights/${bright["_key"]}`}>
                            <div>
                                <h3>@{bright.author.slice(5)}</h3>
                                <p>{truncate(bright.content, 50, true)}</p>
                            </div>
                        </Link>
                    </li>
                {/each}
            </ul>
        </div>
    </div>
</div>

<main>
    {#if brights.length == 0}
        <p>No hay brillos para mostrar..</p>
        <!-- <div> -->
        <img src={Tumbleweed} alt="animacion" />
        <!-- <img src="~/assets/tumbleweed.png" alt=" " /> -->
        <!-- </div> -->
    {:else}
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
                    comments={bright.comments}
                    media={bright.media}
                    rebrightedBy={bright.rebrightedBy?.username}
                />
            {/each}
            <InfiniteLoading
                distance={200}
                on:infinite={(e) => fetchBrights(brights.length, e)}
            />
        </section>
    {/if}

    <div class:active={see}>
        <button on:click={() => (see = false)}><Close /></button>
        <FormBrillo route="/brights" />
    </div>

    <button on:click={() => (see = true)}><PencilPlus /></button>
</main>

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

        section {
            height: 80vh;
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

        :global(img) {
        }
    }

    div.principal {
        display: grid;
        grid-template-columns: 1fr 1fr;
        justify-content: space-evenly;
        justify-items: baseline;
        margin: 32px;
        > div {
            justify-self: end;

            input {
                border-radius: 16px;
                padding: 4px 16px;
            }
            > div {
                position: absolute;
                background-color: var(--background-color);
                z-index: 99;
            }
        }
        :global(button) {
            background: none;
            border: 0;
            // padding: 16px;
        }
    }

    ul {
        list-style: none;
        li {
            border: 1px solid var(--dark-primary-color);
            border-radius: 8px;
            padding: 8px;
            width: 100%;
        }
    }

    button {
        cursor: pointer;
    }

    @keyframes rebota {
        0% {
            transform: translate(0, 0), rotate(0deg);
        }
        30% {
            transform: translate(30vw, 0), rotate(360deg);
        }

        60% {
            transform: translate(50vw, 0), rotate(720deg);
        }
        100% {
            transform: translate(80vw, 0), rotate(1080deg);
        }
    }
</style>
