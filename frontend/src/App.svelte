<script>

    import { Router, Route } from "svelte-routing";
    import Profile from "pages/Profile";
    import Index from "pages/Index";

</script>

<Router>
    
    <Route path="/" component={Index} />
    <Route path="/user/:username" component={Profile}></Route>

</Router>

<style global>
    * {
        margin: 0;
        padding: 0;
    }
</style>