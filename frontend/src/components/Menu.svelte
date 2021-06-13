<script>
    import auth from "utils/auth";
    import Menu from "svelte-material-icons/Menu";
    import { clickOutside } from "utils/strings";
    import { Link, navigate } from "svelte-routing";

    let active = false;
    async function logout() {
        // try {
        await fetch(API_URL + "/logout", { credentials: "include" });
        $auth = null;

        navigate(`/`);
        /* window.location.href = "/"; */
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
    button {
        display: from-lg(none);
        z-index: 1;
    }
    nav {
        height: 60vh;
        width: 10%;
        position: fixed;
        z-index: 1;
        left: -100% from-lg(5%);
        top: 10%;
        background-color: var(--primary-color);
        overflow-x: hidden;
        transition: 0.5s;

        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;

        border-radius: 8px;

        font-weight: bolder;
        font-size: 1.2em;
        padding: 24px;

        &.active {
            left: 0;
        }
        ul {
            list-style: none;
        }

        p {
            cursor: pointer;
        }
    }
</style>
