# Verden
A large scale economic simulation

## TODO

Based on common project requirements, these are the files that you will likely need to adjust:

backend/main.go: Configure CORS for your specific domains and set up proper route handling, including error handling.
backend/Dockerfile: Configure for multi-stage build to have a lean production image.
frontend/Dockerfile: Similarly, set up a multi-stage build for a production-ready frontend image.
docker-compose.yml: Update for any additional services, network definitions, or changes to existing services.
frontend/src/routes/Simulation/+page.svelte: Use environment variables for API endpoints rather than hardcoded URLs.
.gitignore: Make sure it includes all auto-generated files or directories that shouldn't be pushed to the repository.