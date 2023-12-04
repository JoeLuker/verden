<script lang="ts">
    import { onMount } from 'svelte';
    import * as d3 from 'd3';

    let svg: d3.Selection<SVGSVGElement, unknown, HTMLElement, any>;

    onMount(async () => {
        // Create the SVG container
        svg = d3.select<SVGSVGElement, unknown>('#myDiagramDiv')
            .append('svg')
            .attr('width', '100%')
            .attr('height', '400px')
            .style('background-color', '#DAE4E4');

        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/api/get-diagram-id`);
            if (response.ok) {
                const diagramID = await response.text();
                console.log('Diagram ID:', diagramID);
                return diagramID;
            }
        } catch (error) {
            console.error('Error fetching diagram ID:', error);
        }

        // Fetch data and render the diagram
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/api/get-diagram`);
            if (response.ok) {
                const data = await response.json();
                console.log('Data fetched from MongoDB:', data);

                // Assuming data.nodes and data.links exist
                renderDiagram(data.nodes, data.links);
            }
        } catch (error) {
            console.error('Error fetching form structure:', error);
        }
    });



    function renderDiagram(nodes: { x: number, y: number }[], links: { source: { x: number, y: number }, target: { x: number, y: number } }[]) {
        // Render the nodes
        svg.selectAll<SVGCircleElement, unknown>('.node')
            .data(nodes)
            .enter()
            .append('circle')
            .attr('class', 'node')
            .attr('cx', (d: { x: number, y: number }) => d.x)
            .attr('cy', (d: { x: number, y: number }) => d.y)
            .attr('r', 10)
            .style('fill', 'blue');

        // Render the links
        svg.selectAll<SVGLineElement, unknown>('.link')
            .data(links)
            .enter()
            .append('line')
            .attr('class', 'link')
            .attr('x1', (d: { source: { x: number } }) => d.source.x)
            .attr('y1', (d: { source: { y: number } }) => d.source.y)
            .attr('x2', (d: { target: { x: number } }) => d.target.x)
            .attr('y2', (d: { target: { y: number } }) => d.target.y)
            .style('stroke', 'black');
    }
</script>

<div id="myDiagramDiv"></div>
