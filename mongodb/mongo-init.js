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
var diagramStructure = {
    Nodes: [
        { Key: "Dynamic Project Diagram", Color: "red" },
        { Key: "Backend Service", Color: "blue" },
        { Key: "CI/CD Pipeline", Color: "green" },
        { Key: "Infrastructure Automation", Color: "yellow" },
        { Key: "Frontend", Color: "purple" }
    ],
    Links: [
        { From: "Dynamic Project Diagram", To: "Backend Service" },
        { From: "Backend Service", To: "CI/CD Pipeline" },
        { From: "CI/CD Pipeline", To: "Infrastructure Automation" },
        { From: "Infrastructure Automation", To: "Frontend" }
    ]
};

// Insert the Structure into a Collection
newDB.diagramStructures.insert(diagramStructure);
