<script>
    import Input from "components/Input";
    import { Link } from "svelte-routing";
    import SVGheader from "assets/loginHeader.svg";
    import { calculateAge } from "utils/date";
    let username = "",
        usernameError = "";
    let password = "",
        passwordError = "";
    let name = "",
        nameError = "";
    let birthday = "",
        birthdayError = "";
    let email = "",
        emailError = "";

    async function signup() {
        let birthdayDate = new Date(birthday);
        let edad = calculateAge(birthdayDate);

        if (edad < 18) {
            birthdayError = "You need to have more than 18 years";
            return;
        }
        let birthdayms = Date.parse(birthdayDate);

        let res = await fetch(API_URL + "/user", {
            method: "POST",
            body: new URLSearchParams({
                username,
                password,
                name,
                email,
                birthday: birthdayms,
            }),
        });

        let data = await res.json();

        switch (res.status) {
            case 400:
                birthdayError = " Error birthday format";
                return;

            case 409:
                usernameError = "Username already exists";
                //document.getElementById("username").focus
                return;
        }
    }

    async function check_username() {
        usernameError = "";
        // let res = await fetch(API_URL + "/user/login", {
        //     method: "POST",
        //     body: JSON.stringify({
        //         username,
        //         password: "",
        //     }),
        // });

        // usernameError = "";

        // switch (res.status) {
        //     case 401:
        //         usernameError = "User already exist";
        //         return;
        // }

        let res = await fetch(API_URL + "/user/exits", {
            method: "POST",
            body: JSON.stringify({
                username,
                password,
            }),
        });

        switch (res.status) {
            case 404:
                usernameError = "User already exist";
                return;
            case 401:
                passwordError = "You need password";
                return;
        }

        let data = await res.json();

        console.log(data);
    }
</script>

<section>
    {@html SVGheader}

    <h2>Create Account</h2>
    <form on:submit|preventDefault={signup}>
        <div id="id1">
            <Input
                type="text"
                label="Username"
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

            <Link to="/">Before</Link>
            <a href="#id2" on:click|preventDefault={check_username}>Next</a>
        </div>

        <div id="id2">
            <Input
                type="email"
                label="Email"
                id="email"
                bind:value={email}
                errorMessage={emailError}
                regex={/(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))/}
                invalidInputMessage="incorrect email"
            />
            <a href="#id1">Before </a>
            <a href="#id3">Next</a>
        </div>
        <div id="id3">
            <Input
                type="text"
                label="Name"
                id="name"
                bind:value={name}
                errorMessage={nameError}
            />

            <Input
                type="date"
                label="Birthday"
                id="birthday"
                bind:value={birthday}
                bind:errorMessage={birthdayError}
                regex={/^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$/}
                invalidInputMessage="incorrect birthday"
            />

            <!-- ^\d{4}-([0]\d|1[0-2])-([0-2]\d|3[01])$ -->
            <div>
                <a href="#id2">Before</a>
                <input type="submit" value="Sign up" />
            </div>
        </div>
    </form>
</section>

<style lang="scss">
    section {
        // padding: 32px;
        display: flex;
        flex-direction: column;
        // height: 100vh;
        justify-content: space-between;
        h2 {
            padding: 0 32px;
            margin-top: 16px;
        }
        > :global(svg) {
            align-self: flex-start;
            order: from-md(99);
            transform: from-md(scaleY(-1));
            max-height: from-md(30vw);
        }

        > form {
            scroll-snap-type: x mandatory;
            display: flex;
            overflow: hidden;
            width: 100%;
            scroll-behavior: smooth;

            > div {
                scroll-snap-align: center;
                width: 100%;
                flex: 0 0 100%;
                // margin: 16px;
                padding: 16px;

                :global(div) {
                    margin-bottom: 16px;
                    // padding: 0 16px;
                }
                :global(a),
                input[type="submit"] {
                    padding: 8px 16px;
                    text-align: center;
                    border-radius: 12px;
                    font-weight: bold;
                    background-color: var(--primary-color);
                    border: 2px solid var(--primary-color);
                    font-size: 14px;
                }
            }
        }
    }
</style>
