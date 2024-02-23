# Load Kubernetes service definitions
k8s_yaml(['k8s/backend-deployment.yaml', 'k8s/backend-service.yaml', 'k8s/frontend-deployment.yaml', 'k8s/frontend-service.yaml', 'k8s/redis-deployment.yaml', 'k8s/redis-service.yaml'])

# Docker image build instructions
docker_build('backend-image', './backend')
# Define a Docker build for the Svelte frontend
docker_build('frontend-image', './frontend', 
    live_update=[
        fall_back_on('./frontend/vite.config.js'),
        sync('./frontend/src', '/app/src'),
        run(
            'yarn install',
            trigger=['./frontend/package.json', './frontend/yarn.lock']
        )
    ]
)

k8s_resource('frontend-deployment', port_forwards='5735:5173')
