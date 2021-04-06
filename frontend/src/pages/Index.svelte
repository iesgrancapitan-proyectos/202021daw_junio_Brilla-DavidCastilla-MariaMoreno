<script>
    import { Link } from "svelte-routing";
</script>

<section>
    <form>
        <div>
            <input type="text" name="user" id="user" />
            <label for="user">User</label>
        </div>

        <div>
            <input type="password" name="password" id="password" />
            <label for="password">Password</label>
        </div>

        <div class="buttons">
            <Link>Sign up</Link>
            <input type="submit" value="Log in" />
        </div>
    </form>

    <hr />

    <div class="google_facebook">
        <Link class="google">Log in with Google</Link>
        <Link class="facebook">Log in with Facebook</Link>
    </div>
</section>

<style lang="scss">
    section {
        padding: 16px;
    }
    form > div {
        position: relative;
        margin: 16px 0;
        label {
            background-color: white;
            margin: 0 8px;
            position: absolute;
            top: 50%;
            transform: translateY(-50%);
            left: 16px;
            font-weight: bold;
            padding: 0px 2px;
        }

        input {
            padding: 16px;
            display: block;
            width: 100%;
            border: 2px solid black;
            border-radius: 8px;

            &:focus {
                outline: none;
                border: 2px solid var(--primary-color);
                + label {
                    top: 0%;
                    color: var(--dark-primary-color);
                }
            }
        }
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
            filter: drop-shadow(0px 4px 8px rgba(249, 197, 95, 0.3));
        }

        input{
            background-color: var(--primary-color);
            border: 2px solid var(--primary-color);


        }
    }

    .google_facebook {
        display: flex;
        flex-direction: column;
        text-align: center;

        :global(a) {
            color: white;
            margin: 12px;
            padding: 16px;
        }

        :global(.google) {
            background: linear-gradient(91.44deg, #db4437 0%, #dc5429 100%);
            box-shadow: 0px 4px 8px rgba(219, 68, 55, 0.3);
            border-radius: 12px;
        }

        :global(.facebook) {
            background: linear-gradient(91.44deg, #4267b2 0%, #4254b2 100%);
            box-shadow: 0px 4px 8px rgba(66, 103, 178, 0.3);
            border-radius: 12px;
        }
    }
</style>
