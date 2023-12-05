<script lang="ts">
  import { onMount, afterUpdate } from 'svelte';
  import { forceSimulation, forceLink, forceManyBody, forceCenter } from 'd3-force';
  import type { Simulation, SimulationNodeDatum, SimulationLinkDatum, ForceLink } from 'd3';

  export let nodes: SimulationNodeDatum[] = [];
  export let links: SimulationLinkDatum<SimulationNodeDatum>[] = [];

  let simulation: Simulation<SimulationNodeDatum, SimulationLinkDatum<SimulationNodeDatum>>;

  onMount(() => {
    simulation = forceSimulation(nodes)
      .force('link', forceLink(links).id(d => `${d.index ?? 0}`)) // Ensuring the id is a string or number
      .force('charge', forceManyBody())
      .force('center', forceCenter(800 / 2, 600 / 2));

    simulation.on('tick', () => {
      // Update logic on tick
    });
  });

  afterUpdate(() => {
    const linkForce = simulation.force('link') as ForceLink<SimulationNodeDatum, SimulationLinkDatum<SimulationNodeDatum>>;
    if (linkForce) {
      linkForce.links(links);
      simulation.nodes(nodes).alpha(1).restart();
    }
  });
</script>

<svg id="diagram" width="800" height="600">
  <!-- SVG rendering logic goes here -->
</svg>
