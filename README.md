# MyNginx Controller

A Kubernetes controller that manages nginx deployments through custom resources.

## Overview

The MyNginx Controller creates and manages nginx deployments based on `MyNginx` custom resources. It supports configurable replicas, optional ConfigMap integration for custom HTML content, and optional Secret integration for TLS certificates.

## Features

- ✅ **Custom Resource Management**: Define nginx deployments via `MyNginx` CRD
- ✅ **Scalable Deployments**: Configurable replica count
- ✅ **Lifecycle Management**: Proper cleanup with finalizers
- ✅ **ConfigMap Integration**: Mount custom HTML content (planned)
- 🔄 **TLS Support**: Mount TLS certificates from Secrets (planned)
- 🔄 **Status Reporting**: Rich error and status information (planned)

## Quick Start

### Prerequisites
- Kubernetes cluster (1.32+)
- kubectl configured
- Go 1.23.0+ (for development)

### Installation
```bash
# Deploy the controller
make deploy

# Create a sample MyNginx resource
kubectl apply -f examples/basic-mynginx.yaml

# Check the deployment
kubectl get mynginx
kubectl get deployment
```

### Usage Example
```yaml
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: my-nginx
spec:
  replicas: 3
  indexConfigMapName: my-index
```

## Documentation

- 📘 [Claude AI Guide](claude.md) - Comprehensive guide for AI assistants
- 🚀 [Examples](EXAMPLES.md) - Usage examples and scenarios  
- 🛠 [Contributing](CONTRIBUTING.md) - Development guidelines
- 📋 [TODO](TODO.md) - Planned features and tasks
- 🤖 [AI Prompts](AI_PROMPTS.md) - AI assistant prompts and examples
- 🏗 [Architecture Decisions](ADR.md) - Design decisions and rationale

## Requirements (Original Specification)

- The controller should read a CRD with definitions of:
 - Number of replicas (default: 1)
 - Name of ConfigMap to use as content for the file "/usr/share/nginx/html/index.html" (optional)
 - Name of Secret to set TLS (optional)

- Then will create a deployment with NGINX image.

- If ConfigMap or Secret don't exist, update status field with failure message.
- If CRD changes, reconcile.
- If ConfigMap changes, ignore.
- If Secret changes, reconcile.

## Development

### Setup
```bash
# Clone and setup
git clone <repository>
cd mynginx-controller
make install-tools

# Run tests
make test

# Run locally
make run
```

### Common Commands
- `make` - Generate code and build
- `make test` - Run tests
- `make deploy` - Deploy to cluster
- `make undeploy` - Remove from cluster

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## Architecture

```
MyNginx CRD → Controller → Nginx Deployment
     ↓             ↓              ↓
 ConfigMap ────────┴─────────→ Volume Mount
     ↓             ↓              ↓  
 Secret ───────────┴─────────→ TLS Config
```

## Project Status

**Current**: Basic controller with replica and ConfigMap management  
**Next**: Secret integration, status reporting
