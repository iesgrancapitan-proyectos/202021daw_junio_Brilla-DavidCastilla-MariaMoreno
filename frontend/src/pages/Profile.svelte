<script>
let user="";
let name="";
let username="", bio="", followed="", followers="", nbrillos="";
</script>



<section >
    <div>
        <svg height="100" width="100">
            <circle
                cx="200"
                cy="200"
                r="200"
                stroke="#f9c55f"
                stroke-width="3"
                fill="#f9c55f"
            />
        </svg>
    </div>
    <p>{username}</p>

    <img src="/img/{user}" alt="img perfil">

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
    <hr>
</section>

<style lang="scss">
</style>
