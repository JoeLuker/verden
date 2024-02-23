<script lang="ts">
  import InputForm from './InputForm.svelte';
  import SimulationOutput from './SimulationOutput.svelte';

  type SimulationResultType = { EconomyHealth: number; MarketTrends: number; } | null;
  let simulationResult: SimulationResultType = null;

  function handleSimulation(event: CustomEvent<SimulationResultType>) {
  simulationResult = event.detail;
}

</script>

<section>
  <h1>Economic Simulation</h1>
  <InputForm on:simulationComplete={handleSimulation} />
  <SimulationOutput {simulationResult} />
</section>

<style>
  section {
    font-family: 'Arial', sans-serif;
    max-width: 800px;
    margin: auto;
    padding: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }

  h1 {
    color: #333;
    text-align: center;
  }
</style>