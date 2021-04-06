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
    :root{
        --primary-color: #F9C55F ;
        --secondary-color: #FB5666;
        --dark-primary-color: #F7B024;
        --light-primary-color: #FBDA9A;
        --dark-seconday-color: #FA4553;
        --light-seconday-color: #FC8994;
    }

    * {
        margin: 0;
        padding: 0;
    }
</style>