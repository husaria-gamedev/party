<script lang="ts">
    import { userStore } from '$lib/store';

    type UserData = {
        code: string;
        name: string;
    };

    // Something isn't working here, needs a fix
    let userData: UserData
    userStore.subscribe((val) => userData = val)



    
</script>

<div>
    <h1>Lobby</h1>
    {#if userData}
        <h1>Welcome, {userData.name}!</h1>
    {/if}
</div>