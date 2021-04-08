<script>
 import Input from "components/Input";

let username = "",
    usernameError = "";
let password = "",
    passwordError = "";
let name = "",
    nameError = "";


</script>

<form>
    <Input 
    type="text"
    label="Name"
    id="name"
    bind:value={name}
    errorMessage={nameError}
    />
    
    <Input
        type="text"
        label="User"
        id="username"
        bind:value={username}
        errorMessage={usernameError}
    />

    <Input
        type="password"
        label="Password"
        id="password"
        bind:value={password}
        errorMessage={passwordError}
    />

    <div>
        <input type="submit" value="Singup" />
    </div>

</form>


<style lang="scss">

</style>