<script>
    let brilloContent = "";
    let contador = brilloContent.length;
    $: contador = brilloContent.length;
    let file = [];

    async function createBrillo() {
        let form = new FormData();

        const reg_url = /(https?:\/\/)?((\w+\.)?\w{2,}\.\w{2,})/gm;

        brilloContent = brilloContent.replace(
            reg_url,
            `<a href="http://$2">$2<a>`
        );

        form.append("content", brilloContent);

        // for (let index = 0; index < file.length; index++) {
        //     form.append(`media_${index}`, file[index]);
        // }
        form.append("media", file);
        console.log(file);

        let res = await fetch(API_URL + "/brights", {
            method: "POST",
            body: form,
            credentials: "include",
        });

        // location.reload();
    }
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
<section>
    <form on:submit|preventDefault={createBrillo} enctype="multipart/form-data">
        <textarea
            name="brillo"
            id="brillo"
            cols="30"
            rows="10"
            bind:value={brilloContent}
        />

        <input type="file" src="file" alt="file" bind:files={file} multiple />
        <div>
            <span class:red={contador > 250} id="contador">{contador}/250</span>
            <input type="submit" disabled={contador > 250} value="Enviar" />
        </div>
    </form>
</section>

<style lang="scss">
    section {
        margin: 32px;
        text-align: center;
        textarea {
            border: solid 2px #f9c55f;
            resize: none;
            width: 100%;
        }

        div {
            display: flex;
            justify-content: space-between;
            margin-top: 16px;
        }
    }

    .red {
        color: red;
    }
</style>
