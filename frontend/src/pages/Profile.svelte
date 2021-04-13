<script>
    import auth from "utils/auth";
    let user = "";
    let name = "";
    export let username = "",
        bio = "",
        followed = "",
        followers = "",
        nbrillos = "";
</script>

<section>
    <div>
        <svg height="36vw" width="100vw">
            <circle
                cx="50vw"
                cy="-25vw"
                r="60vw"
                stroke="#f9c55f"
                stroke-width="3"
                fill="#f9c55f"
            />
        </svg>
    </div>
    <p>{username}</p>

    <img src="/img/{user}" alt="img perfil" />

    <p>{name}</p>

    <div>
        <p>{bio}</p>

        <div>
            <p>{followed} Followed</p>
            <p>{followers} followers</p>
            <p>{nbrillos} Brillos</p>
        </div>
    </div>
    <button>Follow/Following</button>
    <hr />
</section>

<style lang="scss">
</style>
