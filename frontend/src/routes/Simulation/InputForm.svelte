<script lang="ts">

  let formData: { [key: string]: string } = {
    parameter1: '',
    parameter2: '',
    // Add more parameters as needed
  };
  
  async function submitForm(event: SubmitEvent): Promise<void> {
    event.preventDefault();
      const response = await fetch('/api/simulate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
      const data = await response.json();
      // Handle the response
    }
  </script>
  
  <form on:submit={submitForm}>
    <input type="text" bind:value={formData.parameter1} placeholder="Parameter 1">
    <input type="text" bind:value={formData.parameter2} placeholder="Parameter 2">
    <!-- Add more input fields as needed -->
    <button type="submit">Run Simulation</button>
  </form>
  