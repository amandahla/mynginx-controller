# Architecture Decision Records (ADR)

This document tracks important architectural and design decisions for the MyNginx Controller project.

## ADR-001: Custom Resource Definition Structure

**Date**: 2025-01-05  
**Status**: Current  

### Context
Need to define the API structure for MyNginx custom resources that will be managed by the controller.

### Decision
- Use `webapp.mynginx.amandahla.xyz` as the API group
- Start with `v1` API version
- Keep spec minimal with room for future expansion:
  ```go
  type MyNginxSpec struct {
      Replicas int32 `json:"replicas,omitempty"`
  }
  ```

### Consequences
- Simple, focused API surface
- Easy to extend with ConfigMap and Secret references
- Follows Kubernetes API conventions
- Future changes will require careful versioning

---

## ADR-002: Controller Runtime Framework

**Date**: 2025-01-05  
**Status**: Current  

### Context
Need to choose a framework for implementing the Kubernetes controller.

### Decision
Use controller-runtime v0.20.2 with the following patterns:
- Reconciler interface for business logic
- Manager for controller lifecycle
- Structured logging with context
- Finalizers for cleanup

### Rationale
- Industry standard for Kubernetes controllers
- Good documentation and community support
- Handles common patterns (watches, caching, etc.)
- Compatible with Kubebuilder tooling

### Consequences
- Consistent with other operators
- Good performance and reliability
- Learning curve for new contributors
- Dependency on controller-runtime evolution

---

## ADR-003: Resource Management Strategy

**Date**: 2025-01-05  
**Status**: Current  

### Context
Controller needs to manage nginx deployments and handle resource lifecycle.

### Decision
- One deployment per MyNginx resource
- Use finalizers for cleanup coordination
- Owner references for garbage collection
- Reconcile on MyNginx changes

### Implementation
```go
const MyNginxFinalizer = "webapp.mynginx.amandahla.xyz/finalizer"
```

### Consequences
- Clean resource cleanup
- Predictable resource relationships
- Proper Kubernetes garbage collection
- Additional complexity in finalizer handling

---

## ADR-004: Error Handling and Status Reporting

**Date**: 2025-01-05  
**Status**: Planned  

### Context
Need strategy for handling errors and reporting status to users.

### Decision
Implement status conditions pattern:
```go
type MyNginxStatus struct {
    Conditions []metav1.Condition `json:"conditions,omitempty"`
    ReadyReplicas int32 `json:"readyReplicas,omitempty"`
}
```

### Rationale
- Standard Kubernetes pattern
- Rich error reporting capability
- Enables monitoring and alerting
- Good user experience

### Consequences
- More complex status management
- Better observability
- Requires additional RBAC permissions
- Need proper condition types definition

---

## ADR-005: ConfigMap Integration Design

**Date**: 2025-01-05  
**Status**: Planned  

### Context
Requirements specify optional ConfigMap integration for custom HTML content.

### Decision
- Add `ConfigMapName` field to spec
- Mount as volume to `/usr/share/nginx/html/`
- Validate ConfigMap existence
- Report errors in status
- Do NOT reconcile on ConfigMap changes (per requirements)

### Implementation Plan
```go
type MyNginxSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    ConfigMapName string `json:"configMapName,omitempty"`
}
```

### Consequences
- Users can customize nginx content
- Need additional validation logic
- Status reporting becomes important
- ConfigMap changes won't trigger updates

---

## ADR-006: TLS/Secret Integration Design

**Date**: 2025-01-05  
**Status**: Planned  

### Context
Requirements specify optional Secret integration for TLS certificates.

### Decision
- Add `SecretName` field to spec
- Mount as TLS certificates
- Enable HTTPS port (443)
- Reconcile on Secret changes
- Validate Secret format and existence

### Implementation Plan
```go
type MyNginxSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    ConfigMapName string `json:"configMapName,omitempty"`
    SecretName string `json:"secretName,omitempty"`
}
```

### Consequences
- Supports secure deployments
- More complex deployment configuration
- Need Secret change detection
- Additional validation requirements

---

## ADR-007: Testing Strategy

**Date**: 2025-01-05  
**Status**: Current/Partial  

### Context
Need comprehensive testing strategy for controller reliability.

### Decision
Multi-layer testing approach:
- Unit tests for business logic
- Integration tests with fake clients
- E2E tests with real Kubernetes
- Use Ginkgo/Gomega for consistency

### Current State
- Basic test structure exists
- Need comprehensive coverage
- E2E tests partially implemented

### Consequences
- High confidence in changes
- Slower development initially
- Better long-term maintainability
- Requires test environment setup

---

## ADR-008: Deployment Configuration

**Date**: 2025-01-05  
**Status**: Current  

### Context
Need to decide on nginx deployment configuration details.

### Decision
- Use official nginx image
- Standard nginx configuration
- Configurable replica count
- Proper resource labels
- Owner references for cleanup

### Implementation
```go
// Standard nginx deployment with configurable replicas
// Labels: app=mynginx, mynginx=<resource-name>
// Owner reference to MyNginx resource
```

### Consequences
- Predictable nginx behavior
- Easy to understand and debug
- Limited customization options
- Standard nginx security model

---

## Future ADRs to Consider

- **Observability**: Metrics, tracing, health checks
- **Security**: Pod security context, network policies
- **Performance**: Resource limits, HPA integration
- **Extensibility**: Webhook validation, multiple versions
- **Operations**: Backup, disaster recovery, monitoring
