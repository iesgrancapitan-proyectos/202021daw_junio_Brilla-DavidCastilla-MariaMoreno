<script>
    import FilePond from "svelte-filepond";
    let brilloContent = "";
    let contador = brilloContent.length;
    $: contador = brilloContent.length;
    let file;

    async function createBrillo() {
        let form = new FormData();

        const reg_url = /(https?:\/\/)?((\w+\.)?\w{2,}\.\w{2,})/gm;

        brilloContent = brilloContent.replace(
            reg_url,
            `<a href="http://$2">$2<a>`
        );

        form.append("content", brilloContent);
        let arrFile = file.getFiles().map((l) => l.file);

        for (let i = 0; i < arrFile.length; i++) {
            form.append(`media_${i}`, arrFile[i]);
        }

        let res = await fetch(API_URL + "/brights", {
            method: "POST",
            body: form,
            credentials: "include",
        });

        location.reload();
    }
</script>

<section>
    <form on:submit|preventDefault={createBrillo} enctype="multipart/form-data">
        <textarea
            name="brillo"
            id="brillo"
            cols="30"
            rows="10"
            bind:value={brilloContent}
        />

        <FilePond
            bind:this={file}
            allowMultiple={true}
            credits={null}
            max-files={4}
        />

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
