<script>
    import { Link } from "svelte-routing";
    import Input from "components/Input";

    let username = "",
        usernameError = "";
    let password = "",
        passwordError = "";

    function doLogin() {}
</script>

<section>
    <form on:submit|preventDefault={doLogin}>
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

        <div class="buttons">
            <Link to="/signup">Sign up</Link>
            <input type="submit" value="Log in" />
        </div>
    </form>

    <hr />

    <a class="google">Log in with Google</a>
    <a class="facebook">Log in with Facebook</a>
</section>

<style lang="scss">
    section {
        padding: 32px;
        display: grid;
        grid-gap: 16px;

        form {
            display: grid;
            grid-gap: 16px;
        }

        .buttons {
            display: flex;
            justify-content: space-between;

            :global(a),
            input {
                padding: 8px;
                flex-basis: 30%;
                text-align: center;
                border-radius: 12px;
                font-weight: bold;
            }

            :global(a) {
                color: var(--secondary-color);
                border: 2px solid var(--secondary-color);
                box-shadow: 0px 4px 8px rgba(251, 86, 102, 0.6);
            }

            input {
                background-color: var(--primary-color);
                border: 2px solid var(--primary-color);
            }
        }

        hr {
            margin: 16px 0;
        }

        > :global(a) {
            color: white;
            font-weight: bold;
            text-align: center;
            padding: 16px;

            &.google {
                background: linear-gradient(91.44deg, #db4437 0%, #dc5429 100%);
                box-shadow: 0px 4px 8px rgba(219, 68, 55, 0.3);
                border-radius: 12px;
            }

            &.facebook {
                background: linear-gradient(91.44deg, #4267b2 0%, #4254b2 100%);
                box-shadow: 0px 4px 8px rgba(66, 103, 178, 0.3);
                border-radius: 12px;
            }
        }
    }
</style>
