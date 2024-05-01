# Interview Challenge Implementation

## Overview
This project implements a REST API based on the provided OpenAPI/Swagger definition. The API is built using Golang and provides several endpoints as specified in the challenge requirements.

## Features

### Endpoints
1. **Root Endpoint (`/`):**
   - Provides the current date in UNIX epoch format.
   - Returns the service version.
   - Includes a boolean property `Kubernetes` indicating if the application is running under Kubernetes.

2. **Health Endpoint (`/health`):**
   - Provides information about the health of the service.

3. **Metrics Endpoint (`/metrics`):**
   - Exposes Prometheus metrics for monitoring and instrumentation.

4. **Lookup Endpoint (`/v1/tools/lookup`):**
   - Resolves IPv4 addresses for a given domain.
   - Logs successful queries and their results in a PostgreSQL database.

5. **Validate Endpoint (`/v1/tools/validate`):**
   - Validates if the input is an IPv4 address.

6. **History Endpoint (`/v1/history`):**
   - Retrieves the latest 20 saved queries from the database, ordered by the most recent first.

### Dockerized Development Environment
- Fully Dockerized development environment using Docker and Docker Compose.
- Services and tools are included in the Docker Compose definition.
- Supports easy setup and initialization with `docker-compose up -d --build`.

### DevOps Tasks
- [x] **Kubernetes Support:** Includes Kubernetes manifests or a Helm Chart for deploying the application in Kubernetes.
- [x] **Secrets Management:** Sensitive data such as database passwords stored in Kubernetes secrets.
- [x] **CI Pipeline:** Utilizes GitHub Actions for continuous integration tasks:
  - Basic tests and linting.
  - Build and package the application into a versioned Docker image.
  - Generate Helm chart as an artifact for deployment.

### CI/CD Pipeline
- GitHub Actions CI pipeline is configured in .github/workflows/ci.yml.
- Automatically triggered upon each commit or pull request
- Executes tests, linting, builds Docker images, and generates Helm chart artifacts.

### Kubernetes Deployment
- Helm Chart provided in the `f2pool-chart/` directory.
- Secrets management for sensitive data using Opaque secret (created using base64 encoding)
- Supports deployment to a Kubernetes cluster for production-ready setups.

## Setup
```sh
# Build and up docker
docker-compose up --build -d

# Helm package
docker compose build && docker compose push
helm install f2pool f2pool-chart

# Check if service is running
kubectl get po
kubectl get svc
minikube tunnel
```