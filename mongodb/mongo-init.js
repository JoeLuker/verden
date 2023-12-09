// Admin Authentication
db.getSiblingDB('admin').auth(
    process.env.MONGO_INITDB_ROOT_USERNAME,
    process.env.MONGO_INITDB_ROOT_PASSWORD
);

// User Creation
db.createUser({
    user: process.env.MONGO_APP_USERNAME,
    pwd: process.env.MONGO_APP_PASSWORD,
    roles: [{role: "readWrite", db: "diagramDB"}]
});

// Switch to New Database (creates if does not exist)
var newDB = db.getSiblingDB('diagramDB');

// Create the Structure for DiagramNode, DiagramLink, and DiagramStructure
// Updated to reflect D3's SimulationNodeDatum and SimulationLinkDatum types
var diagramStructure = {
    Nodes: [
        { key: "Dynamic Project Diagram", fx: 100, fy: 100, color: "red" },
        { key: "Backend Service", fx: 200, fy: 200, color: "blue" },
        { key: "CI/CD Pipeline", fx: 300, fy: 300, color: "green" },
        { key: "Infrastructure Automation", fx: 400, fy: 400, color: "yellow" },
        { key: "Frontend", fx: 500, fy: 500, color: "purple" }
    ],
    Links: [
        { source: "Dynamic Project Diagram", target: "Backend Service" },
        { source: "Backend Service", target: "CI/CD Pipeline" },
        { source: "CI/CD Pipeline", target: "Infrastructure Automation" },
        { source: "Infrastructure Automation", target: "Frontend" }
    ]
};

// Insert the Structure into a Collection
newDB.diagramStructures.insert(diagramStructure);
