<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte';
  import type { FormField, FormStructure, FormData, Category } from '$lib/types/formTypes';

  const dispatch = createEventDispatcher();
  let formStructure: FormStructure = {};
  let formData: FormData = {};




  onMount(async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/form-structure`);
      if (response.ok) {
        const data = await response.json();
        // Assuming data.Categories is an array of { CategoryName, Fields }
        data.Categories.forEach((category: Category) => {
          category.Fields.forEach((field: FormField) => {
            formData[field.name] = field.defaultValue;
            formStructure[field.name] = field;
          });
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
        const errorText = await response.text();
        throw new Error(`Failed to run simulation: ${errorText}`);
      }

      const simulationResults = await response.json();
      dispatch('simulationComplete', simulationResults);
    } catch (error: unknown) {
      if (error instanceof Error) {
        console.error('Error:', error.message);
        alert(error.message);
      }
    }
  }
</script>

