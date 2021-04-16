<script>
    import { Router, Route } from "svelte-routing";
    import Profile from "pages/Profile";
    import Index from "pages/Index";
    import Signup from "pages/Signup";
    import Timeline from "pages/Timeline";
    import auth from "utils/auth";
    import { onMount } from "svelte";

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
    <Route path="/activate" component={Profile} />
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
