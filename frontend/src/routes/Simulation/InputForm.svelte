<script lang="ts">

  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher();

  let formData: { parameter1: number, parameter2: number } = {
    parameter1: 0,
    parameter2: 0,
    // Add more parameters as needed
  };

  async function submitForm(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/api/simulate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        // Handle non-OK responses from the server
        const errorText = await response.text();
        console.error('Error from server:', errorText);
        return;
      }

      const data = await response.json();

      console.log('Received data:', data);
      // simulationResult = data;
      const simulationResults = data; // Replace with your actual results

      dispatch('simulationComplete', simulationResults);

      // Handle the response
      // Update your UI or state with the received data
    } catch (error) {
      console.error('Fetch error:', error);
      // Handle fetch errors (e.g., network issues)
    }
  }
</script>

<form on:submit={submitForm}>
  <input type="number" bind:value={formData.parameter1} placeholder="Parameter 1">
  <input type="number" bind:value={formData.parameter2} placeholder="Parameter 2">
  <!-- Add more input fields as needed -->
  <button type="submit">Run Simulation</button>
</form>
