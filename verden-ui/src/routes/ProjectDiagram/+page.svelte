<script lang="ts">
  import ProjectDiagram from './ProjectDiagram.svelte';
  import { onMount } from 'svelte';
  import type { SimulationNodeDatum, SimulationLinkDatum } from 'd3';
  import { fetchDiagramData } from './services/apiService';

  let error: string = '';
  let nodes: SimulationNodeDatum[] = [];
  let links: SimulationLinkDatum<SimulationNodeDatum>[] = [];

  onMount(async () => {
    try {
      const diagramData = await fetchDiagramData();
      nodes = diagramData.nodes;
      links = diagramData.links;
    } catch (e) {
      console.error('Error loading diagram:', e);
      error = 'Unable to load the diagram. Please try again later.';
    }
  });
</script>

<main>
  {#if error}
    <p class="error-message">{error}</p>
  {:else}
    <ProjectDiagram {nodes} {links} />
  {/if}
</main>

<style>
  .error-message {
    color: red;
    /* Additional styling for error message */
  }
</style>
