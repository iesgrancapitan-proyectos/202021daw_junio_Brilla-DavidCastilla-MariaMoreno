<script>
    import { fly } from "svelte/transition";
    import Eye from "svelte-material-icons/Eye";
    import EyeOff from "svelte-material-icons/EyeOff";

    export let type;
    export let label;
    export let value;
    export let errorMessage;
    export let regex = /.*/;
    export let invalidInputMessage = "Input not valid";

    let inputRef;
    let originalType = type;

    $: errorMessage && inputRef && value ? inputRef.focus() : void 0;

    $: regex.test(value)
        ? (errorMessage = "")
        : (errorMessage = invalidInputMessage);

    /**
     * @param {Event} e
     */
    function setValue(e) {
        switch (type) {
            case "number":
            case "range":
                value = toNumber(e.target.value);
            case "file":
                value = e.target.files;
            default:
                value = e.target.value;
        }
    }

    const changeVisibility = () => {
        type = type === "password" ? "text" : "password";
    };
</script>

<div>
    <input
        {...$$props}
        {type}
        class:not_empty={value}
        class:invalid={errorMessage && value}
        id={$$props.id}
        name={$$props.id}
        bind:this={inputRef}
        on:change={(type === "file" || type === "range") && setValue}
        on:input={type !== "file" && setValue}
        {value}
    />
    <label for={$$props.id}>{label}</label>
    {#if originalType === "password"}
        <div on:click={changeVisibility}>
            <svelte:component this={type === "password" ? Eye : EyeOff} />
        </div>
    {/if}
    {#if errorMessage && value}
        <p transition:fly={{ duration: 440, x: -500 }}>{errorMessage}</p>
    {/if}
</div>

<style lang="scss">
    div {
        position: relative;

        label {
            position: absolute;
            transform: translateY(16px);
            top: 0px;
            left: 0px;
            margin: 0 16px;
            background-color: var(--background-color);
            transition: all 440ms;
        }

        div {
            position: absolute;
            display: block;
            right: 0;
            padding: 16px;
            font-size: 1.25rem;
            top: 0;
            cursor: pointer;

            > :global(svg) {
                display: block;
            }
        }

        input {
            border-radius: 12px;
            padding: 16px;
            border: 1px solid black;
            width: 100%;

            &:focus,
            &.not_empty {
                outline: none;
                border-color: var(--primary-color);

                + label {
                    transform: translateY(-50%);
                    padding: 0 2px;
                    color: var(--primary-color);
                }
            }

            &.invalid {
                border-color: red;

                + label {
                    color: red;
                }
            }
        }

        p {
            font-size: 0.9rem;
            color: red;
            margin: 4px 12px;
        }
    }
</style>
