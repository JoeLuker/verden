<script lang="ts">
    import InputForm from './InputForm.svelte';
    import SimulationOutput from './SimulationOutput.svelte';
  
    let simulationResult: any = null;  // Define a more specific type if possible
  
    // Function to handle the response from InputForm
    function handleSimulation(data: any): void {  // Adjust the type as needed
        simulationResult = data;
    }
  </script>
  
  <section>
    <h1>Economic Simulation</h1>
    <InputForm on:simulationComplete={handleSimulation} />
    <SimulationOutput {simulationResult} />
  </section>
  