<script lang="ts">
  import ProjectDiagram from './ProjectDiagram.svelte';

  let key = "";
  let value = "";

  async function createNewNode() {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/create-diagram-node`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ key, value }),
      });
  
      if (response.ok) {
        console.log("Value saved to MongoDB");
      }
    }

</script>

  
 <main>
    <h1> Project Digram </h1>
    <input type="text" bind:value={key} placeholder="Key" />
    <input type="text" bind:value={value} placeholder="Value to store" />
    <button on:click={createNewNode}>Save to Redis</button>

    <ProjectDiagram />
</main>

<style>
    /* Add your CSS here */
</style>