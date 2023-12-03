<script lang="ts">
    import { onMount } from 'svelte';
    import * as go from 'gojs';

    let myDiagram: go.Diagram;

    onMount(async () => {
        const $ = go.GraphObject.make;

        // Define the model (nodes and links)
        myDiagram.model = $(go.GraphLinksModel, {
            linkKeyProperty: 'key'  // this should be defined according to your data structure
        });

        // Fetch data and set it to the model
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/api/diagram`);
            if (response.ok) {
                const data = await response.json();
                console.log('Data fetched from MongoDB:', data);

                // Assuming data.nodes and data.links exist
                myDiagram.model.nodeDataArray = data.nodes;
                myDiagram.model.linkDataArray = data.links;
            }
        } catch (error) {
            console.error('Error fetching form structure:', error);
        }
    });
</script>


<div id="myDiagramDiv" style="width:100%; height:400px; background-color: #DAE4E4;"></div>
