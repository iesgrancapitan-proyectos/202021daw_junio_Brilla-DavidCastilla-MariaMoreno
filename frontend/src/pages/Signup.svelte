<script>
    import Input from "components/Input";

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
        let res = await fetch(API_URL + "/user/", {
            method: "POST",
            body: new URLSearchParams({
                username,
                password,
                name,
                email,
                birthday,
            }),
        });
        console.log(res);
        let data = await res.text();
        console.log(data);
    }
</script>

<section>
    <h2>Create Account</h2>
    <form on:submit|preventDefault={signup}>
        <Input
            type="text"
            label="Name"
            id="name"
            bind:value={name}
            errorMessage={nameError}
        />

        <Input
            type="email"
            label="Email"
            id="email"
            bind:value={email}
            errorMessage={emailError}
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

        <Input
            type="text"
            label="Birthday"
            id="birthday"
            bind:value={birthday}
            errorMessage={birthdayError}
        />

        <div>
            <input type="submit" value="Sign up" />
        </div>
    </form>
</section>

<style lang="scss">
    section {
        padding: 32px;
    }

    form {
        display: grid;
        grid-gap: 16px;
        input {
            padding: 8px 16px;
            text-align: center;
            border-radius: 12px;
            font-weight: bold;
            background-color: var(--primary-color);
            border: 2px solid var(--primary-color);
        }
    }

    h2 {
        text-align: center;
        margin-bottom: 32px;
    }
</style>
