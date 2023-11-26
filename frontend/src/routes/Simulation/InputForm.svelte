<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte';
  import type { FormField, FormStructure, FormData } from '$lib/types/formTypes';

  interface FormField {
    type: string;
    defaultValue: number | string;
    placeholder: string;
  }

  interface FormStructure {
    [key: string]: FormField;
  }

  interface FormData {
    [key: string]: number | string;
  }


  const dispatch = createEventDispatcher();
  let formStructure: FormStructure = {}; // Use the FormStructure interface
  let formData: FormData = {}; // Use the FormData interface

  onMount(async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/form-structure`);
      if (response.ok) {
        const data = await response.json();
        // Process each field and add it to formData and formStructure
        data.fields.forEach((field: any) => {
          formData[field.name] = field.defaultValue;
          formStructure[field.name] = field;
        });
      }
    } catch (error) {
      console.error('Error fetching form structure:', error);
    }
  });



  async function submitForm(event: SubmitEvent) {
    event.preventDefault();
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/simulate`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        const errorText = await response.json();
        console.error('Error from server:', errorText.message);
        alert(`Failed to run simulation: ${errorText.message}`);
        return;
      }

      const simulationResults = await response.json();
      dispatch('simulationComplete', simulationResults);
    } catch (error) {
      console.error('Network error:', error);
      alert("Failed to run simulation due to a network error");
    }
  }
</script>

<form on:submit={submitForm}>
  {#each Object.entries(formStructure) as [key, value]}
    {#if value.type === 'number'}
      <input type="number" bind:value={formData[key]} placeholder={value.placeholder} />
    {:else if value.type === 'text'}
      <input type="text" bind:value={formData[key]} placeholder={value.placeholder} />
    {/if}
    <!-- Add more conditions for other types if needed -->
  {/each}
  <button type="submit">Run Simulation</button>
</form>

