<script>
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import auth from "utils/auth";

    let token = new URL(document.location).searchParams.get("token");

    onMount(async () => {
        let res = await fetch(API_URL + "/user/activate", {
            method: "POST",
            body: JSON.stringify({ token }),
        });

        console.log(res);

        let { username } = await res.json();
        $auth = username;

        navigate("/");
    });
</script>
