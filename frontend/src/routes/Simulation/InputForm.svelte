<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte';
  import type { FormField, FormStructure, FormData, Category } from '$lib/types/formTypes';

  const dispatch = createEventDispatcher();
  let formStructure: FormStructure = { categories: [] };
  let formData: FormData = {};

  onMount(async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/form-structure`);
      if (response.ok) {
        const data = await response.json();
        formStructure.categories = data.categories;
        formStructure.categories.forEach((category: Category) => {
          formData[category.categoryName] = {}; // Create a nested object for each category
          category.fields.forEach((field: FormField) => {
            formData[category.categoryName][field.name] = field.defaultValue;
          });
        });
      }
    } catch (error) {
      console.error('Error fetching form structure:', error);
    }
  });

  async function submitForm(event: SubmitEvent) {
    event.preventDefault();
    console.log(formData)
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

<form on:submit|preventDefault={submitForm}>
  {#each formStructure.categories as category}
    <fieldset>
      <legend>{category.categoryName}</legend>
      {#each category.fields as field}
        <div class="form-group">
          <label for={`${category.categoryName}-${field.name}`}>{field.name}</label>
          <input id={`${category.categoryName}-${field.name}`}
                 bind:value={formData[category.categoryName][field.name]}
                 type="number" />
        </div>
      {/each}
    </fieldset>
  {/each}
  <button type="submit">Submit</button>
</form>

<style>
  form {
  display: grid;
  grid-gap: 20px;
  margin-top: 20px;
}

fieldset {
  border: 1px solid #ddd;
  padding: 10px;
  border-radius: 5px;
}

legend {
  padding: 0 10px;
  font-weight: bold;
  color: #444;
}

.form-group {
  display: grid;
  grid-template-columns: 1fr 2fr;
  align-items: center;
  grid-gap: 10px;
}

.form-group label {
  text-align: right;
  font-size: 14px;
}

.form-group input[type="number"] {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button[type="submit"] {
  width: 100%;
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button[type="submit"]:hover {
  background-color: #45a049;
}

</style>