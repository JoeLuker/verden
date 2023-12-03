<script lang="ts">
  export let simulationResult: { EconomyHealth: number; MarketTrends: number; } | null = null;

  function formatSimulationResult(result: { EconomyHealth: number; MarketTrends: number; }): string {
    return `Economy Heatlth: ${result.EconomyHealth.toFixed(2)}, Market Trends: ${result.MarketTrends.toFixed(2)}`;
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

<style>
  div {
  background-color: #f2f2f2;
  padding: 15px;
  border-radius: 5px;
  margin-top: 20px;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
}

h2 {
  color: #4CAF50;
  margin-bottom: 10px;
}

p {
  font-size: 16px;
  color: #333;
}

</style>