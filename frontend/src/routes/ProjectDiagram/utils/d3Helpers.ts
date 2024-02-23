import * as d3 from 'd3';

export const drawDiagram = (
    nodes: d3.SimulationNodeDatum[],
    links: d3.SimulationLinkDatum<d3.SimulationNodeDatum>[]
) => {
    const svg = d3.select<SVGSVGElement, unknown>('#diagram'); // Specify the element types
    const width = +svg.attr("width");
    const height = +svg.attr("height");

    svg.selectAll("*").remove();

    const zoom = d3.zoom<SVGSVGElement, unknown>()
        .scaleExtent([1 / 2, 8])
        .on("zoom", event => {
            g.attr("transform", event.transform);
        });

    const g = svg.append("g");
    svg.call(zoom); // Apply zoom to the SVG

    const simulation = d3.forceSimulation(nodes)
        .force("link", d3.forceLink(links).id(d => (d as d3.SimulationNodeDatum).index?.toString() ?? ''))
        .force("charge", d3.forceManyBody())
        .force("center", d3.forceCenter(width / 2, height / 2));

    const link = g.append("g")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.6)
        .selectAll("line")
        .data(links)
        .join("line")
        .attr("x1", d => (d.source as d3.SimulationNodeDatum).x ?? 0)
        .attr("y1", d => (d.source as d3.SimulationNodeDatum).y ?? 0)
        .attr("x2", d => (d.target as d3.SimulationNodeDatum).x ?? 0)
        .attr("y2", d => (d.target as d3.SimulationNodeDatum).y ?? 0);

    // Node rendering
    const node = g.append("g")
      .selectAll<SVGCircleElement, d3.SimulationNodeDatum>("circle")
      .data(nodes)
      .join("circle")
      .attr("r", 20)
      .attr("fill", "blue");

      node.call(drag(simulation));

    // Drag functionality
    function drag(simulation: d3.Simulation<d3.SimulationNodeDatum, undefined>) {
        function dragstarted(event: d3.D3DragEvent<SVGCircleElement, d3.SimulationNodeDatum, d3.SimulationNodeDatum>) {
            if (!event.active) simulation.alphaTarget(0.3).restart();
            event.subject.fx = event.subject.x;
            event.subject.fy = event.subject.y;
        }

        function dragged(event: d3.D3DragEvent<SVGCircleElement, d3.SimulationNodeDatum, d3.SimulationNodeDatum>) {
            event.subject.fx = event.x;
            event.subject.fy = event.y;
        }

        function dragended(event: d3.D3DragEvent<SVGCircleElement, d3.SimulationNodeDatum, d3.SimulationNodeDatum>) {
            if (!event.active) simulation.alphaTarget(0);
            event.subject.fx = null;
            event.subject.fy = null;
        }

        return d3.drag<SVGCircleElement, d3.SimulationNodeDatum>()
            .on("start", dragstarted)
            .on("drag", dragged)
            .on("end", dragended);
    }

    // Tooltip setup
    const tooltip = d3.select("body").append("div")
        .attr("class", "tooltip")
        .style("position", "absolute")
        .style("visibility", "hidden")
        .style("background-color", "white")
        .style("border", "solid")
        .style("border-width", "1px")
        .style("border-radius", "5px")
        .style("padding", "10px");

    node.on("mouseover", (event, d) => {
        tooltip.html(`Node: ${d.index}`)
            .style("visibility", "visible")
            .style("left", (event.pageX + 10) + "px")
            .style("top", (event.pageY - 10) + "px");
    })
    .on("mouseout", () => {
        tooltip.style("visibility", "hidden");
    });

    // Update simulation on every tick
    simulation.on("tick", () => {
        link
            .attr("x1", d => (d.source as d3.SimulationNodeDatum).x ?? 0)
            .attr("y1", d => (d.source as d3.SimulationNodeDatum).y ?? 0)
            .attr("x2", d => (d.target as d3.SimulationNodeDatum).x ?? 0)
            .attr("y2", d => (d.target as d3.SimulationNodeDatum).y ?? 0);

        node
            .attr("cx", d => d.x ?? 0)
            .attr("cy", d => d.y ?? 0);
    });
};
