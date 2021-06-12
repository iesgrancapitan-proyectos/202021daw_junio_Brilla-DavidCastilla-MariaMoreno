<script>
    import auth from "utils/auth";
    import Menu from "svelte-material-icons/Menu";
    import { clickOutside } from "utils/strings";
    import { Link } from "svelte-routing";

    let active = false;
    async function logout() {
        // try {
        await fetch(API_URL + "/logout", { credentials: "include" });
        $auth = null;

        // navigate(`/timeline`);
        window.location.href = "/";
        // } catch (e) {}
    }
</script>

<button on:click={() => (active = !active)}><Menu /></button>
<nav class:active use:clickOutside on:click_outside={() => (active = false)}>
    <ul>
        <Link to={"/"}><li>Home</li></Link>
        <Link to={`/user/${$auth}`}><li>Profile</li></Link>
    </ul>

    <p on:click={logout}>Logout</p>
</nav>

<style lang="scss">
    nav {
        //ocultar menu
        height: 60vh;
        width: 0;
        position: fixed;
        z-index: 1;
        background-color: var(--primary-color);
        overflow-x: hidden;
        transition: 0.5s;

        //@posicion
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        border-radius: 16px;

        font-weight: bolder;
        font-size: 1.2em;

        &.active {
            width: 60vw;
            padding: 24px;
        }
        ul {
            list-style: none;
        }

        p {
            cursor: pointer;
        }
    }
</style>
