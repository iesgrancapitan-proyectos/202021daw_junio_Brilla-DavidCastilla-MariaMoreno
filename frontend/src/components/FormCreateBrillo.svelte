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

        for (let i = 0; i < file.length; i++) {
            form.append(`media_${i}`, file[i]);
        }

        let res = await fetch(API_URL + "/brights", {
            method: "POST",
            body: form,
            credentials: "include",
        });

        location.reload();
    }

    async function verimg() {
        // Creamos el objeto de la clase FileReader
        let reader = new FileReader();

        // Leemos el archivo subido y se lo pasamos a nuestro fileReader
        reader.readAsDataURL(file);

        // Le decimos que cuando este listo ejecute el código interno
        reader.onload = function () {
            let preview = document.getElementById("preview"),
                image = document.createElement("img");

            image.src = reader.result;
            image.style.width = "50px";
            image.style.height = "50px";

            preview.innerHTML = "";
            preview.append(image);
        };
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

        <input
            type="file"
            src="file"
            alt="file"
            id="file"
            bind:files={file}
            multiple
            on:change={verimg}
        />
        <div id="preview" />

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
