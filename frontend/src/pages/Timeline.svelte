<script>
    import { Router, Route } from "svelte-routing";
    import Profile from "pages/Profile";
    import Index from "pages/Index";
    import Signup from "pages/Signup";
    import Timeline from "pages/Timeline";
    import Activate from "pages/Activate";
    import auth from "utils/auth";
    import { onMount } from "svelte";
    import { Carousel } from "svelte-images";
    const { Modal } = Carousel;

    onMount(async () => {
        try {
            let res = await fetch(API_URL + "/refresh", {
                credentials: "include",
            });
            if (res.status != 200) return;

            let { username } = await res.json();
            $auth = username;
        } catch (e) {
            console.log(e);
        }
    });
</script>

<Modal />
<Router>
    <Route path="/">
        {#if !$auth}
            <Index />
        {:else}
            <Timeline />
        {/if}
    </Route>
    <Route path="/signup" component={Signup} />
    <Route path="/user/:username" component={Profile} />
    <Route path="/activate" component={Activate} />
</Router>

<style global>
    :root {
        --primary-color: #f9c55f;
        --secondary-color: #fb5666;
        --dark-primary-color: #f7b024;
        --light-primary-color: #fbda9a;
        --dark-seconday-color: #fa4553;
        --light-seconday-color: #fc8994;

        --background-color: white;
    }

    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: "Red Hat Text", "Noto Sans JP", sans-serif;
    }

    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
        font-family: "Red Hat Display", "Noto Sans JP", sans-serif;
    }

    a {
        text-decoration: none;
        color: black;
    }
</style>
                {rebrillos}
            </button>
            <button on:click|preventDefault={null}
                ><CommentMultipleOutline />
                {comments}
            </button>
            <button on:click|preventDefault={null}>
                <StarOutline />
                {interactions}
            </button>
        </div>

        <span>{humanDate.relativeTime(uploadDate)}</span>
    </div>
</article>

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

        .carousel {
            scroll-snap-type: x mandatory;
            display: flex;
            overflow: hidden;
            width: 100%;
            scroll-behavior: smooth;
            img {
                scroll-snap-align: center;
                width: 100%;
                flex: 0 0 100%;
            }
        }
    }

    h1 {
        font-size: 16px;
    }
    h2 {
        font-size: 14px;
    }
</style>
