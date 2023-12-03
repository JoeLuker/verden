<script lang="ts">
    import { onMount } from 'svelte';
    import * as go from 'gojs';

    let myDiagram: go.Diagram;

    onMount(() => {
        const $ = go.GraphObject.make;

        myDiagram = $(go.Diagram, "myDiagramDiv", {
            initialContentAlignment: go.Spot.Center,
            "undoManager.isEnabled": true
        });

        myDiagram.nodeTemplate =
            $(go.Node, "Auto",
                $(go.Shape, "RoundedRectangle", { strokeWidth: 0 },
                    new go.Binding("fill", "color")),
                $(go.TextBlock, { margin: 8 },
                    new go.Binding("text", "key"))
            );

        // Define the model (nodes and links)
        myDiagram.model = new go.GraphLinksModel(
            [
                { key: "Dynamic Project Diagram", color: "lightblue" },
                { key: "Backend Service", color: "orange" },
                { key: "CI/CD Pipeline", color: "lightgreen" },
                { key: "Infrastructure Automation", color: "pink" },
                { key: "Frontend", color: "yellow" },
                // Add other components as necessary
            ],
            [
                { from: "Dynamic Project Diagram", to: "Backend Service" },
                { from: "Backend Service", to: "CI/CD Pipeline" },
                { from: "CI/CD Pipeline", to: "Infrastructure Automation" },
                { from: "Infrastructure Automation", to: "Frontend" },
                // Define other relationships as necessary
            ]);
    });
</script>

<div id="myDiagramDiv" style="width:100%; height:400px; background-color: #DAE4E4;"></div>
