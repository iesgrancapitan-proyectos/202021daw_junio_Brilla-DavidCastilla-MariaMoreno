<script>
    import { Link } from "svelte-routing";
</script>

<section>

    <form>
        <div>
            <input type="text" name="user" />
            <label for="user">User</label>
        </div>
    
        <div>
            <input type="password" name="password" id="password" />
            <label for="password">Password</label>
        </div>
    
        <div class="buttons">
            <Link>Sign up</Link>
            <input type="submit" value="Log in">
        </div>
        
    </form>
    
    <hr>
    
    <div class="google_facebook">
        <Link class="google">Log in with Google</Link>
        <Link class="facebook">Log in with Facebook</Link>
    </div>
 
    
</section>

<style lang="scss">
    section{
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
        }

        input {
            padding: 16px;
            display: block;
            width: 100%;
            &:focus {
                outline: none;
                border-color: var(--primary-color);
                border-radius: 8px;
                + label {
                    top: 0%;
            
                }
            }
        }
    }

    .buttons{
        display: flex;
        justify-content: space-between;
        
        :global(a), input{
            padding: 8px;
            flex-basis: 30%;
            
        }
    }

    .google_facebook{
        display: flex;
        flex-direction: column;
        text-align: center;
        :global(a){
            padding: 8px;
                        
        }
    }
</style>
