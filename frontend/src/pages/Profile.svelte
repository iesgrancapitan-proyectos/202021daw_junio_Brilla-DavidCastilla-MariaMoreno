<script>
    import { onMount } from "svelte";
    import circulo from "assets/circulo.svg";
    import auth from "utils/auth";

    let name = "";
    export let username = "";
    let bio = "",
        followed = "",
        followers = "",
        nbrillos = "",
        imgPerfil = "";

    onMount(async () => {
        let res = await fetch(API_URL + `/user/${username}`);
        let info = await res.json();
        console.log(info);
        bio = info.bio;
        name = info.name;
        imgPerfil = info.imgPerfil;
        imgPerfil = "https://picsum.photos/200";
        //n brillos suma de brillos.
        nbrillos = await fetch(API_URL + `/user/${username}/brights/count`);
        let nbrillosJson = await nbrillos.json();
        nbrillos = nbrillosJson["nbrillos"];

        //calculo
        followed = await fetch(API_URL + `/nfollowed/${username}`);
        let followedJson = await followed.json();
        followed = followedJson["followed"];
        //calculo
        followers = await fetch(API_URL + `/nfollowers/${username}`);
        let followersJson = await followers.json();
        followers = followersJson["followers"];
    });

    // console.log(res);
    //  }
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

        <p>@{username}</p>

        <img src={imgPerfil} alt="img perfil" />

        <p>{name}</p>
    </div>

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
