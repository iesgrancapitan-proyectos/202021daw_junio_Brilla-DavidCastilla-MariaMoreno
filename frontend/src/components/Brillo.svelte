<script>
    import { Link } from "svelte-routing";
    import RepeatVariant from "svelte-material-icons/TwitterRetweet";
    import StarOutline from "svelte-material-icons/StarOutline";
    import CommentMultipleOutline from "svelte-material-icons/CommentMultipleOutline";

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
    <div>
        <Link to={`/user/${user.username}`}>
            <h1>{user.name}</h1>
        </Link>
        <Link to={`/user/${user.username}`}>
            <h2>@{user.username}</h2>
        </Link>
    </div>

    <p>{content}</p>

    <div>
        <div>
            <!-- iconos -->
            <RepeatVariant />
            <CommentMultipleOutline />
            <StarOutline />
        </div>

        <span>{humanDate.relativeTime(uploadDate)}</span>
    </div>
</article>

<style lang="scss">
    article {
        position: relative;
        padding: 16px;
        border: 1px solid var(--primary-color);
        margin-top: 72px * 0.5;
        padding-top: 36px;
        border-radius: 16px;
        img {
            border-radius: 15%;
            position: absolute;
            top: 0;
            left: 50%;
            transform: translate(-50%, -50%);
            size: 72px;
        }

        > div {
            text-align: center;
        }
        > div:last-child {
            display: flex;
            justify-content: space-between;
            margin-top: 16px;
        }
    }

    h1 {
        font-size: 16px;
    }
    h2 {
        font-size: 14px;
    }
</style>
