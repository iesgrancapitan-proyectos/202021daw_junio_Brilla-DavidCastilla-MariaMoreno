<script>
    import { Link } from "svelte-routing";
    export let user;
    export let content;
    export let likes;
    export let comments;
    export let rebrights;
    export let uploadDate;

    import humanDate from "human-date";
</script>

<article>
    <Link to={`/user/${user.username}`}>
        <img
            src={user.img ?? "https://via.placeholder.com/150"}
            alt={`Image profile of user ${user.username}`}
        />
    </Link>
    <Link to={`/user/${user.username}`}>
        <h1>{user.name}</h1>
    </Link>
    <Link to={`/user/${user.username}`}>
        <h2>{user.username}</h2>
    </Link>
    <p>{content}</p>
    <span>{humanDate.relativeTime(uploadDate)}</span>
</article>

<style lang="scss">
    article {
        position: relative;
        padding: 16px;
        border: 1px solid red;
        margin-top: 72px * 0.5;
        img {
            position: absolute;
            top: 0;
            left: 50%;
            transform: translate(-50%, -50%);
            size: 72px;
        }
    }
</style>
