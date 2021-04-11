<script>
    import { Link } from "svelte-routing";
    import loginHeaderSVG from "assets/loginHeader.svg";
    import googleSVG from "assets/google.svg";
    import facebookSVG from "assets/facebook.svg";
    import Input from "components/Input";
    import auth from "utils/auth";

    let username = "",
        usernameError = "";
    let password = "",
        passwordError = "";

    async function doLogin() {
        let res = await fetch(API_URL + "/user/login", {
            method: "POST",
            body: new URLSearchParams({
                username,
                password,
            }),
            credentials: "include",
        });

        usernameError = "";
        passwordError = "";

        switch (res.status) {
            case 404:
                usernameError = "User not found";
                return;
            case 401:
                passwordError = "Pasword incorrect";
                return;
        }

        let tkn = res.headers.get("x-token");
        $auth = tkn;
    }
</script>

<section>
    {@html loginHeaderSVG}
    <form on:submit|preventDefault={doLogin}>
        <header>
            <h1>Enter to Brilla!</h1>
            <div>
                <button class="google">{@html googleSVG}</button>
                <button class="facebook">{@html facebookSVG}</button>
            </div>
        </header>
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
            <Link to="/signup">Sign up</Link>
            <input type="submit" value="Log in" />
        </div>
    </form>
    <div />
</section>

<style lang="scss">
    section {
        display: flex;
        flex-direction: column;
        height: 100vh;
        justify-content: space-between;

        form {
            display: grid;
            padding: 0px 32px;
            grid-gap: 16px;

            header {
                display: flex;
                justify-content: space-between;
                align-items: end;
                > div {
                    display: flex;
                    align-items: end;
                    button {
                        border: none;
                        background-color: red;
                        size: 48px;
                        border-radius: 8px;
                        margin-left: 8px;
                        &.google {
                            background: linear-gradient(
                                45deg,
                                #db4437,
                                #dc5429
                            );
                            box-shadow: 0px 4px 8px rgba(219, 68, 55, 0.3);
                        }

                        &.facebook {
                            background: linear-gradient(
                                45deg,
                                #4267b2,
                                #4254b2
                            );
                            box-shadow: 0px 4px 8px rgba(66, 103, 178, 0.3);
                        }
                        :global(svg) {
                            fill: white;
                            margin: 8px;
                        }
                    }
                }
            }

            > div {
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
                    box-shadow: 0px 4px 8px rgba(251, 86, 102, 0.3);
                }

                input {
                    background-color: var(--primary-color);
                    border: none;
                    box-shadow: 0 4px 8px rgba(249, 197, 95, 0.3);
                }
            }
        }

        > :global(a) {
            color: white;
            font-weight: bold;
            text-align: center;
            padding: 16px;
        }
    }
</style>
