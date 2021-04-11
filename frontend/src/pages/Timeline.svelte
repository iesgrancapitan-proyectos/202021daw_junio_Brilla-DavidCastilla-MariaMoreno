<script>
    import auth from "utils/auth";

    function logout() {
        $auth = null;
    }
</script>

<h1>Timeline</h1>

<button on:click={logout}>Logout</button>
