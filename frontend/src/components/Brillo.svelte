<script>
    export let user;
    export let content;
    export let likes;
    export let comments;
    export let rebrights;
    export let uploadDate;

    import humanDate from "human-date";

</script>

<article>
    <img src={user.img} alt={`Image profile of user ${user.username}`}>
    <h1>{user.name}</h1>
    <h2>{user.username}</h2>
    <p>{content}</p>
    <span>{humanDate.relativeTime(uploadDate)}</span>
</article>

<style>
    
</style>