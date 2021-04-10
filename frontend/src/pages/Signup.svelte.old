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
        console.log(birthday);
        birthdaySplit = birthday.split("-");

        birthday = birthday.UTC(
            birthdaySplit[0],
            birthdaySplit[1] - 1,
            birthdaySplit[2]
        );

        console.log(birthday);

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

        switch (res.status) {
            case 400:
                birthdayError = " Error sintaxis birthday";
                return;

            case 409:
                usernameError = "Username already exists";
                return;
        }
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
            regex={/^\d{4}-([0]\d|1[0-2])-([0-2]\d|3[01])$/}
            invalidInputMessage="incorrect email"
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
            type="date"
            label="Birthday"
            id="birthday"
            bind:value={birthday}
            errorMessage={birthdayError}
            regex={/(?:(?:31(\/|-|\.)(?:0?[13578]|1[02]))\1|(?:(?:29|30)(\/|-|\.)(?:0?[13-9]|1[0-2])\2))(?:(?:1[6-9]|[2-9]\d)?\d{2})$|^(?:29(\/|-|\.)0?2\3(?:(?:(?:1[6-9]|[2-9]\d)?(?:0[48]|[2468][048]|[13579][26])|(?:(?:16|[2468][048]|[3579][26])00))))$|^(?:0?[1-9]|1\d|2[0-8])(\/|-|\.)(?:(?:0?[1-9])|(?:1[0-2]))\4(?:(?:1[6-9]|[2-9]\d)?\d{2})/}
            invalidInputMessage="incorrect birthday"
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
