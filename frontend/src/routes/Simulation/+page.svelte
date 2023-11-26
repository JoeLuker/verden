<script lang="ts">
  import InputForm from './InputForm.svelte';
  import SimulationOutput from './SimulationOutput.svelte';

  type SimulationResultType = { Result1: number; Result2: number; } | null;
  let simulationResult: SimulationResultType = null;

  function handleSimulation(event: CustomEvent<SimulationResultType>) {
  console.log('Simulation complete event received:', event.detail);
  simulationResult = event.detail;
}

</script>

<section>
  <h1>Economic Simulation</h1>
  <InputForm on:simulationComplete={handleSimulation} />
  <SimulationOutput {simulationResult} />
</section>
