<script lang="ts">
    export let simulationResult: any = null;  // Define a more specific type if possible
  </script>
  
  {#if simulationResult}
    <div>
      <h2>Simulation Results</h2>
      <p>{simulationResult}</p>
      <!-- Display results in a structured format -->
    </div>
  {:else}
    <p>No results to display</p>
  {/if}
  