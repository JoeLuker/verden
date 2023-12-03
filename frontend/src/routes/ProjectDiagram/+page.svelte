<script lang="ts">
  import ProjectDiagram from './ProjectDiagram.svelte';

  let key = "";
  let color = "";

  let from = "";
  let to = "";

  async function createNewNode() {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/create-diagram-node`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ key, color }),
      });
  
      if (response.ok) {
        console.log("Value saved to MongoDB");
      }
    }

    async function createNewLink() {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/create-diagram-link`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ from, to }),
      });
  
      if (response.ok) {
        console.log("Value saved to MongoDB");
      }
    }

</script>

  
 <main>
    <h1> Project Digram </h1>
    <div id="DiagramInput">
        <input type="text" bind:value={key} placeholder="Key" />
        <input type="text" bind:value={color} placeholder="Color to store" />
        <button on:click={createNewNode}>Save Node to MongoDB</button>
        <br>
        <input type="text" bind:value={from} placeholder="from" />
        <input type="text" bind:value={to} placeholder="to" />
        <button on:click={createNewLink}>Save Link to MongoDB</button>
    </div>
    <ProjectDiagram />
</main>

<style>
    #DiagramInput {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    /* Add your CSS here */
</style>