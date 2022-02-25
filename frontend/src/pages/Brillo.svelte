<script>
    import { onMount } from "svelte";
    import Brillo from "components/Brillo";
    import ArrowLeft from "svelte-material-icons/ArrowLeft";
    import Menu from "../components/Menu.svelte";
    import { Link } from "svelte-routing";
    import Magnify from "svelte-material-icons/Magnify";
    import debounce from "lodash/debounce";

    export let id;

    let bright;
    let comments;
    let userSearch = [];
    let brightSearch = [];

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
        let data = await fetch(API_URL + `/brights/${id}`);
        bright = await data.json();
        console.log(bright);
    });
</script>

<header>
    <Menu />
    <button on:click={() => history.back()}> <ArrowLeft /></button>
    <h1>Brillo</h1>
    <div>
        <span>
            <input on:input={handleInput} placeholder="Search" />
            <Magnify />
        </span>

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
                                <h3>@{bright.username}</h3>
                                <p>{truncate(bright.content, 50, true)}</p>
                            </div>
                        </Link>
                    </li>
                {/each}
            </ul>
        </div>
    </div>
</header>
<section>
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
    header {
        display: grid;
        grid-template-columns: min-content min-content 1fr min-content from-lg(
                min-content 1fr min-content
            );
        justify-content: space-evenly;
        align-items: center;
        padding: 16px 24px from-md(16px 10%) from-lg(16px 20%);
        border-bottom: 1px solid var(--primary-color);
        box-shadow: 0px 4px 8px rgba(black, 0.2);

        h1 {
            font-size: 1.3rem;
        }

        > div {
            justify-self: end;
            align-self: end;
            > span {
                display: flex;
                align-items: center;
            }
            input {
                outline: none;
                border-radius: 16px;
                border: 1px solid black;
                padding: 4px 16px;
                &:hover,
                &:focus {
                    border-color: var(--secondary-color);
                }
            }
            > div {
                position: absolute;
                background-color: var(--background-color);
                z-index: 99;
            }
            p {
                font-size: 1.5em;
                font-weight: bolder;
            }
        }
        :global(button) {
            font-size: 1.5em;
            background: none;
            border: 0;
            padding: 8px 16px;
            align-self: center;
            height: 42px;
        }
        ul {
            list-style: none;
            li {
                border: 1px solid var(--dark-primary-color);
                border-bottom: 0px;
                border-radius: 0px;
                padding: 8px;
                width: 100%;
                &:last-child {
                    border-radius: 0px 0px 8px 8px;
                    border: 1px solid var(--dark-primary-color);
                }
                &:first-child {
                    border-radius: 8px 8px 0px 0px;
                }
            }
        }
    }
    section {
        width: md-lg(80%) from-lg(60%);
        margin: 32px from-md(32px auto);
        button {
            background: none;
            border: 0;
            font-size: 1.5rem;
            cursor: pointer;
        }
    }
</style>
