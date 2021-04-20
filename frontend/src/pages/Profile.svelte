<script>
    import { onMount } from "svelte";
    import auth from "utils/auth";
    import Input from "components/Brillo";
    import Brillo from "../components/Brillo.svelte";

    let name = "";
    export let username = "";
    let bio = "",
        followed = "",
        followers = "",
        nbrillos = "",
        imgPerfil = "";

    let brillos = "";

    onMount(async () => {
        let res = await fetch(API_URL + `/user/${username}`);
        let info = await res.json();

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

        brillos = await fetch(API_URL + `/user/${username}/brights`);
        brillos = await brillos.json();
    });

    // console.log(res);
    //  }
</script>

<section>
    <div class="inicio">
        <svg height="60vw" width="100vw">
            <circle
                cx="50vw"
                cy="-25vw"
                r="60vw"
                stroke="#f9c55f"
                stroke-width="3"
                fill="#f9c55f"
            />
        </svg>
        <div>
            <p>@{username}</p>

            <img src={imgPerfil} alt="img perfil" />

            <p>{name}</p>
        </div>
    </div>

    <div />

    <div class="info">
        <p>{bio}</p>

        <div>
            <p>{followed} Followed</p>
            <p>{followers} followers</p>
            <p>{nbrillos} Brillos</p>
        </div>

        <button>Follow/Following</button>
        <hr />
    </div>
</section>

<section>
    {#each brillos as brilloo}
        <Brillo
            user={{
                username: brilloo.username,
                name: brilloo.name,
                profile_img: brilloo.profile_img,
            }}
            content={brilloo.content}
            uploadDate={new Date(brilloo.created_at)}
        />
    {/each}
</section>

<style lang="scss">
    section {
        .inicio {
            text-align: center;
            width: 100%;

            svg {
                position: relative;
                z-index: -1;
                top: 0px;
                left: 0px;
            }
            div {
                position: absolute;
                z-index: 0;
                top: 8vh;
                margin: 0 auto;
                width: 100%;
                img {
                    border-radius: 15%;
                    width: 30%;
                }
            }
        }
        .info {
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            justify-content: space-around;
            p {
                flex-basis: 68%;
            }

            div {
                flex-basis: 30%;
            }
            hr {
                flex-basis: 100%;
            }
        }
    }
</style>
