<script>
    import { Link, navigate } from "svelte-routing";
    import RepeatVariant from "svelte-material-icons/TwitterRetweet";
    import StarOutline from "svelte-material-icons/StarOutline";
    import CommentMultipleOutline from "svelte-material-icons/CommentMultipleOutline";
    import humanDate from "human-date";
    import Popover from "svelte-popover";
    import { Carousel } from "svelte-images";
    import { onMount } from "svelte";
    const { open } = Carousel;
    import FormBrillo from "components/FormCreateBrillo";

    import Close from "svelte-material-icons/Close";

    export let user;
    export let content;
    export let interactions;
    export let comments;
    export let rebrillos;
    export let uploadDate;
    export let id;
    export let media = [];
    // let carousel = false;
    let see = false;

    onMount(() => {
        media = media.map((src) => ({ src }));
    });

    const popModal = (idx) =>
        setTimeout(() => {
            open(media, idx);
        }, 0);

    async function rebrillo() {
        let form = new URLSearchParams();
        form.append("brilloId", id);

        let res_rebrilla = await fetch(API_URL + "/brights/rebrilla", {
            method: "POST",
            body: form,
            credentials: "include",
        });
        let { inserted } = await res_rebrilla.json();
        inserted ? rebrillos++ : rebrillos--;
    }

    function newInteraction(type) {
        return async () => {
            let form = new URLSearchParams();
            form.append("brilloId", id);
            form.append("type", type);

            let res = await fetch(API_URL + "/brights/interaction", {
                method: "POST",
                body: form,
                credentials: "include",
            });

            let { inserted } = await res.json();

            inserted ? interactions++ : interactions--;
        };
    }
</script>

<article on:click={() => navigate(`/brights/${id}`)}>
    <Link to={`/user/${user.username}`}>
        <img
            src={user.img ?? "https://via.placeholder.com/150"}
            alt={`Image profile of user ${user.username}`}
        />
    </Link>
    <div>
        <Link to={`/user/${user.username}`}>
            <h1>{user.name}</h1>
        </Link>
        <Link to={`/user/${user.username}`}>
            <h2>@{user.username}</h2>
        </Link>
    </div>

    <p>{@html content}</p>

    {#if media}
        <div class="img">
            {#each media as img, i}
                <img src={img.src} alt="img" on:click={() => popModal(i)} />
            {/each}
        </div>
    {/if}

    <div>
        <div>
            <button on:click|preventDefault={rebrillo}>
                <RepeatVariant />
                <p>{rebrillos}</p>
            </button>
            <button on:click={() => (see = true)}>
                <CommentMultipleOutline />
                <p>{comments}</p>
            </button>
            <Popover
                overlayColor="transparent"
                placement="top-center"
                arrowColor="lightgray"
            >
                <button slot="target">
                    <StarOutline />
                    {interactions}
                </button>
                <div class="content" slot="content">
                    <button on:click={newInteraction("happy")}>Happy</button>
                    <button on:click={newInteraction("cool")}>Cool</button>
                    <button on:click={newInteraction("sad")}>Sad</button>
                    <button on:click={newInteraction("angry")}>Angry</button>
                </div>
            </Popover>
        </div>

        <span>{humanDate.relativeTime(uploadDate)}</span>
    </div>
</article>

<div class="form" class:active={see}>
    <button on:click={() => (see = false)}><Close /></button>
    <FormBrillo route="/brights/comment" brilloKey={id} />
</div>

<style lang="scss">
    article {
        position: relative;
        padding: 16px;
        border: 1px solid var(--primary-color);
        margin-top: 72px * 0.5;
        padding-top: 36px;
        border-radius: 16px;
        :global(a img) {
            border-radius: 15%;
            position: absolute;
            top: 0;
            left: 50%;
            transform: translate(-50%, -50%);
            size: 72px;
        }

        > div {
            text-align: center;
        }
        > div:last-child {
            display: flex;
            justify-content: space-between;
            margin-top: 16px;

            > div {
                display: flex;
            }
        }

        .img {
            display: flex;
            flex-wrap: wrap;

            img {
                width: 50%;
                flex-grow: 1;
                border: 3px solid white;
                height: 150px;
                object-fit: cover;
            }
        }
        button {
            border: 0;
            background: var(--background-color);
            margin-right: 16px;
            display: inline-flex;
            align-items: center;
        }
    }

    .content {
        display: flex;
        padding: 8px 16px;
        background-color: lightgray;
        border-radius: 100vmax;
        box-shadow: 0px 0px 8px 0px rgba(black, 0.5);

        button {
            background-color: red;
            border: 1px solid black;
            height: 48px;
            width: 48px;
            cursor: pointer;
        }
    }

    h1 {
        font-size: 16px;
    }
    h2 {
        font-size: 14px;
    }

    .form {
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
</style>
