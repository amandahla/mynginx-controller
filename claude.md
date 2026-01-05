# Claude AI Assistant Guide - MyNginx Controller

## Project Overview

This is a Kubernetes controller built with Go that manages nginx deployments through Custom Resource Definitions (CRDs). The controller creates and manages nginx pods based on MyNginx custom resources.

## Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   MyNginx CRD   │───▶│  Controller     │───▶│ Nginx Deployment│
│                 │    │  (Reconciler)   │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
        │                       │                       │
        │              ┌─────────────────┐              │
        └─────────────▶│   ConfigMap     │              │
                       │   (HTML content)│              │
                       └─────────────────┘              │
        │              ┌─────────────────┐              │
        └─────────────▶│    Secret       │──────────────┘
                       │  (TLS config)   │
                       └─────────────────┘
```

## Key Components

### 1. Custom Resource Definition (CRD)
- **Location**: [api/v1/mynginx_types.go](api/v1/mynginx_types.go)
- **Purpose**: Defines the MyNginx resource schema
- **Current Spec Fields**:
  - `replicas` (int32): Number of nginx pod replicas

### 2. Controller Logic
- **Location**: [internal/controller/mynginx_controller.go](internal/controller/mynginx_controller.go)
- **Purpose**: Reconciles MyNginx resources to desired state
- **Key Functions**:
  - `Reconcile()`: Main reconciliation loop
  - `reconcileDeployment()`: Ensures nginx deployment exists
  - `reconcileDelete()`: Handles resource cleanup with finalizers

### 3. Main Entry Point
- **Location**: [cmd/main.go](cmd/main.go)
- **Purpose**: Sets up and starts the controller manager

## Development Workflows

### Common Commands
```bash
# Generate code and manifests
make

# Run tests
make test

# Build the controller
make build

# Run locally (requires kubeconfig)
make run

# Deploy to cluster
make deploy

# Remove from cluster
make undeploy
```

### File Structure Patterns
- `api/v1/`: API definitions and types
- `config/`: Kubernetes manifests and kustomization
- `internal/controller/`: Controller implementation
- `test/`: Test files (unit and e2e)

## AI Assistant Guidelines

### When helping with this project:

1. **Code Generation**: Use Go best practices for Kubernetes controllers
2. **CRD Updates**: Remember to run `make` after modifying [api/v1/mynginx_types.go](api/v1/mynginx_types.go)
3. **RBAC**: Update controller comments for new permissions
4. **Testing**: Consider both unit tests and e2e test implications
5. **Versioning**: This is v1 API - breaking changes need new version

### Common Tasks and Patterns

#### Adding a new field to the CRD:
1. Update `MyNginxSpec` in [api/v1/mynginx_types.go](api/v1/mynginx_types.go)
2. Run `make` to regenerate code
3. Update controller logic in [internal/controller/mynginx_controller.go](internal/controller/mynginx_controller.go)
4. Add tests in [internal/controller/mynginx_controller_test.go](internal/controller/mynginx_controller_test.go)

#### Implementing the missing requirements:
The README indicates these features are planned but not yet implemented:
- ConfigMap reference for custom HTML content
- Secret reference for TLS configuration  
- Status field updates for error conditions
- Proper reconciliation on Secret changes (currently ignored for ConfigMap)

### Development Environment Setup
```bash
# Install required tools (if not present)
make install-tools

# Setup test environment
make envtest

# Run with verbose logging
make run ARGS="--log-level=debug"
```

### Debugging Tips
- Controller logs are structured - use `--log-level=debug` for detailed output
- Check CRD generation: `kubectl describe crd mynginxes.webapp.mynginx.amandahla.xyz`
- Inspect controller RBAC: `kubectl describe clusterrole mynginx-controller-role`

## Recent Changes
- **2025-09-19**: Refactored finalizer handling and ensureDeployment logic
- **2025-09-01**: Initial controller implementation with basic deployment management

## Next Development Priorities
1. Implement ConfigMap integration for custom HTML content
2. Add Secret integration for TLS configuration
3. Implement proper status reporting
4. Add comprehensive error handling
5. Improve test coverage

## Dependencies
- Kubernetes 1.32+
- controller-runtime v0.20.2
- Go 1.23.0+
