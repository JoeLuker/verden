<script lang="ts">
  import DiagramCanvas from './components/DiagramCanvas.svelte';
  import type { SimulationNodeDatum, SimulationLinkDatum } from 'd3';

  export let nodes: SimulationNodeDatum[] = [];
  export let links: SimulationLinkDatum<SimulationNodeDatum>[] = [];

  function addNode(node: SimulationNodeDatum) {
    nodes = [...nodes, node];
  }

  function addLink(link: SimulationLinkDatum<SimulationNodeDatum>) {
    links = [...links, link];
  }

  // Implement other functions as needed
</script>

<main>
  <DiagramCanvas {nodes} {links} />
</main>
