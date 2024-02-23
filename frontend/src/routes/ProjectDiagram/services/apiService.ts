import type { SimulationNodeDatum, SimulationLinkDatum } from 'd3';

const API_URL = import.meta.env.VITE_API_URL;

// Utility function to handle fetch requests
async function makeFetchRequest<T>(url: string, options?: RequestInit): Promise<T> {
    const response = await fetch(url, options);
    if (!response.ok) {
        throw new Error(`Error in fetch request: ${response.statusText}`);
    }
    return response.json();
}

export const fetchDiagramData = async (): Promise<{ nodes: SimulationNodeDatum[], links: SimulationLinkDatum<SimulationNodeDatum>[] }> => {
    return makeFetchRequest(`${API_URL}/api/get-diagram`);
};

export const createNode = async (node: SimulationNodeDatum): Promise<SimulationNodeDatum> => {
    return makeFetchRequest(`${API_URL}/api/create-diagram-node`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(node)
    });
};

export const createLink = async (link: SimulationLinkDatum<SimulationNodeDatum>): Promise<SimulationLinkDatum<SimulationNodeDatum>> => {
    return makeFetchRequest(`${API_URL}/api/create-diagram-link`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(link)
    });
};
