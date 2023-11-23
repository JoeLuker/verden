<script lang="ts">
    import { onMount } from 'svelte';

    let simulationParams = { param1: '', param2: '' }; // Define parameters with default values
    let simulationResults: any = null;
    let errorMessage: string = '';

    // Environment variable for the API URL
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';

    async function startSimulation() {
        try {
            const response = await fetch(`${apiUrl}/start-simulation`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(simulationParams),
            });

            if (response.ok) {
                simulationResults = await response.json();
            } else {
                throw new Error("Failed to start simulation");
            }
        } catch (error) {
            if (error instanceof Error) {
                errorMessage = error.message;
            } else {
                errorMessage = 'An unexpected error occurred';
            }
        }
    }

    onMount(() => {
        // Initialization logic if needed
    });
</script>


<main>
    <h1>Economic Simulation</h1>
    <input type="text" bind:value={simulationParams.param1} placeholder="Parameter 1" />
    <input type="text" bind:value={simulationParams.param2} placeholder="Parameter 2" />
    <button on:click={startSimulation}>Start Simulation</button>

    {#if simulationResults}
        <div class="results">
            <pre>{JSON.stringify(simulationResults, null, 2)}</pre>
        </div>
    {/if}

    {#if errorMessage}
        <div class="error-message">{errorMessage}</div>
    {/if}
</main>


<style>
    main {
        font-family: Arial, sans-serif;
        max-width: 600px;
        margin: auto;
        padding: 20px;
        text-align: center;
    }

    input {
        margin: 10px 0;
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    button {
        padding: 10px 15px;
        background-color: #007bff;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    button:hover {
        background-color: #0056b3;
    }

    .results {
        margin-top: 20px;
        background-color: #f8f9fa;
        border: 1px solid #ddd;
        padding: 10px;
    }

    .error-message {
        color: red;
        margin-top: 10px;
    }
</style>
