// Existing code for admin authentication
db.getSiblingDB('admin').auth(
    process.env.MONGO_INITDB_ROOT_USERNAME,
    process.env.MONGO_INITDB_ROOT_PASSWORD
);

// Existing user creation code
db.createUser({
    user: process.env.MONGO_APP_USERNAME,
    pwd: process.env.MONGO_APP_PASSWORD,
    roles: ["readWrite"],
});

// Switch to a new database (it will be created if it does not exist)
var newDB = db.getSiblingDB('diagramDB');

// Create the structure for DiagramNode, DiagramLink, and DiagramStructure
var diagramStructure = {
    ID: new ObjectId(), // MongoDB shell can directly use ObjectId
    Nodes: [
        { Key: "Dynamic Project Diagram", Color: "red" }, // Example, set your own color
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

// Insert the structure into a collection
newDB.diagramStructures.insert(diagramStructure);
