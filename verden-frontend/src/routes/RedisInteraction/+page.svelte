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
      const response = await fetch('http://localhost:8080/redis-set', {
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
      const response = await fetch(`http://localhost:8080/redis-get?key=${key}`);
      const data = await response.json();
      retrievedValue = data.value;
    }

    async function startSimulation() {
        // Implement the logic to start the simulation
        // This might include making a fetch request to your Go backend
        // and handling the response
        // Example (you need to replace the URL and body as per your backend API):
        const response = await fetch('http://localhost:8080/start-simulation', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(simulationParams),
        });

        if (response.ok) {
            simulationResults = await response.json();
            console.log("Simulation started", simulationResults);
        } else {
            console.error("Failed to start simulation");
        }
    }

    function updateParameter(paramName: string, value: any) {
        simulationParams[paramName] = value;
    }
  </script>
  

  
  <main>

    <div>
      <input type="text" bind:value={key} placeholder="Key" />
      <input type="text" bind:value={value} placeholder="Value to store" />
      <button on:click={saveToRedis}>Save to Redis</button>
    
      <button on:click={getFromRedis}>Retrieve from Redis</button>
      {#if retrievedValue}
        <p>Retrieved Value: {retrievedValue}</p>
      {/if}
    </div>

    <h1>Economic Simulation</h1>
    <form on:submit|preventDefault={startSimulation}>
        <!-- Add form fields to capture simulation parameters -->
        <input type="number" bind:value={simulationParams.someNumericParam} />
        <button type="submit">Start Simulation</button>
    </form>

    {#if simulationResults}
        <div>
            <!-- Display simulation results here -->
            <!-- You can use JSON.stringify(simulationResults, null, 2) for raw JSON display -->
        </div>
    {/if}
</main>

<style>
    /* Add your CSS here */
</style>