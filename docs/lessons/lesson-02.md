# Goal

Initialize a Kubernetes controller project using Kubebuilder. This lesson sets up the complete scaffolding needed to build, test, and deploy a controller including build configurations, RBAC, monitoring, and CI workflows.

# Concepts

**Kubebuilder**: A framework for building Kubernetes APIs using Custom Resource Definitions (CRDs). It generates boilerplate code and project structure following Kubernetes best practices.

**controller-runtime**: A Go library that provides building blocks for controllers. It handles watch mechanisms, reconciliation queues, caching, and client interactions with the Kubernetes API server.

**RBAC (Role-Based Access Control)**: Kubernetes authorization mechanism that defines what operations a service account can perform. Controllers need RBAC rules to read and modify cluster resources.

**Prometheus**: A monitoring system that collects metrics. Kubebuilder scaffolds metric endpoints so controllers can expose operational data like reconciliation counts and errors.

**Kustomize**: A template-free configuration management tool for Kubernetes. It uses overlays to customize YAML manifests without modifying the original files.

**Leader Election**: A mechanism that ensures only one controller instance is active when running multiple replicas. This prevents conflicting reconciliation operations.

# Result

This lesson initialized the project with Kubebuilder, creating:
- A `cmd/main.go` entry point with controller manager setup and leader election
- Configuration directories (`config/`) with Kustomize manifests for RBAC, metrics, monitoring, and network policies
- A Dockerfile for building controller container images using multi-stage builds and distroless base images
- E2E test scaffolding in `test/e2e/` and utilities in `test/utils/`
- CI workflows (`.github/workflows/`) for linting, unit tests, and e2e tests
- Updated Makefile with targets for generating manifests, running tests, building images, and deploying
- Go dependencies including `controller-runtime`, `client-go`, and testing frameworks

# References

- [Kubebuilder Book](https://book.kubebuilder.io/)
- [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)
- [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- [Kustomize](https://kustomize.io/)
- [Leader Election](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection)
