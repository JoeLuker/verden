<script lang="ts">
    import { onMount, createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();
    let key = "";
    let value = "";
    let retrievedValue = "";
    type SimulationParams = {
        [key: string]: any;  // This allows any string as a key and any type as a value
        someNumericParam: number;  // You can still define known properties
    };

    let simulationParams: SimulationParams = {
        someNumericParam: 0,
    };
    let simulationResults: any = null; // To store simulation results

    async function saveToRedis() {
      const response = await fetch('http://localhost:8080/api/redis-set', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ key, value }),
      });
  
      if (response.ok) {
        console.log("Value saved to Redis");
      }
    }
  
    async function getFromRedis() {
      const response = await fetch(`http://localhost:8080/api/redis-get?key=${key}`);
      const data = await response.json();
      retrievedValue = data.value;
    }
  </script>
  

  
  <main>
    <h1> Redis Test </h1>
    <div>
      <input type="text" bind:value={key} placeholder="Key" />
      <input type="text" bind:value={value} placeholder="Value to store" />
      <button on:click={saveToRedis}>Save to Redis</button>
    
      <button on:click={getFromRedis}>Retrieve from Redis</button>
      {#if retrievedValue}
        <p>Retrieved Value: {retrievedValue}</p>
      {/if}
    </div>
</main>

<style>
    /* Add your CSS here */
</style>