<script lang="ts">
  export let simulationResult: { Result1: number; Result2: number; } | null = null;

  // Define the type for the parameter
  function formatSimulationResult(result: { Result1: number; Result2: number; }): string {
    return `Result 1: ${result.Result1}, Result 2: ${result.Result2}`;
  }
</script>

{#if simulationResult}
  <div>
    <h2>Simulation Results</h2>
    <p>{formatSimulationResult(simulationResult)}</p>
  </div>
{:else}
  <p>No results to display</p>
{/if}
