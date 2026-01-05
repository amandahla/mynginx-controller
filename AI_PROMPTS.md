# AI Assistant Prompts & Examples

This file contains useful prompts and examples for AI assistants working with this Kubernetes controller project.

## Common AI Assistant Requests

### 1. Implementing Missing Requirements

**Prompt**: "Help me implement ConfigMap integration for the MyNginx controller. The controller should accept a ConfigMapName in the spec and mount it as /usr/share/nginx/html/index.html in the nginx deployment. Update the status field if the ConfigMap doesn't exist."

**Context to provide**:
- Current CRD definition in [api/v1/mynginx_types.go](api/v1/mynginx_types.go)
- Controller logic in [internal/controller/mynginx_controller.go](internal/controller/mynginx_controller.go)
- Requirements from [README.md](README.md)

### 2. Adding TLS Support

**Prompt**: "Add Secret integration for TLS certificates. The MyNginx spec should accept a SecretName field, and the controller should mount the secret and configure nginx for HTTPS. The controller should reconcile when the Secret changes."

**Files to reference**:
- Current deployment creation logic
- RBAC permissions that may need updating
- Status reporting requirements

### 3. Improving Error Handling

**Prompt**: "Improve error handling in the MyNginx controller. Add proper status reporting with conditions, better error categorization, and structured logging with more context."

**Key considerations**:
- Kubernetes status conditions pattern
- controller-runtime error handling best practices
- Existing finalizer implementation

### 4. Adding Tests

**Prompt**: "Help me write comprehensive tests for the MyNginx controller, including unit tests for the reconciliation logic and integration tests for the complete workflow."

**Testing context**:
- Current test structure in [internal/controller/mynginx_controller_test.go](internal/controller/mynginx_controller_test.go)
- Ginkgo/Gomega patterns
- Test environment setup

### 5. Code Review and Refactoring

**Prompt**: "Review the MyNginx controller code for best practices, potential issues, and refactoring opportunities. Focus on controller-runtime patterns and Kubernetes operator best practices."

**Review areas**:
- Reconciliation logic
- Error handling
- Resource management
- RBAC permissions
- Code organization

## Useful Code Snippets for AI Context

### Current CRD Structure
```go
type MyNginxSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    // Missing: ConfigMapName, SecretName fields
}

type MyNginxStatus struct {
    // Missing: status fields for error reporting
}
```

### Controller Pattern
```go
func (r *MyNginxReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // 1. Fetch MyNginx resource
    // 2. Handle finalizers
    // 3. Handle deletion
    // 4. Reconcile deployment
    // Missing: ConfigMap/Secret validation, status updates
}
```

### Expected Deployment Structure
The controller should create nginx deployments with:
- Configurable replica count
- Optional ConfigMap mount for HTML content
- Optional Secret mount for TLS certificates
- Proper labels and owner references

## Development Context for AI

### Project Architecture
- **Pattern**: Kubernetes Operator using controller-runtime
- **Language**: Go 1.23
- **Dependencies**: controller-runtime v0.20.2, Kubernetes 1.32+
- **Build System**: Make-based with code generation

### Key Development Rules
1. Always run `make` after modifying CRD types
2. Update RBAC annotations when adding new permissions
3. Use structured logging with request context
4. Implement proper finalizer cleanup
5. Follow controller-runtime error handling patterns

### Common Gotchas
- Generated code in `zz_generated.deepcopy.go` - don't edit manually
- RBAC permissions need to be complete for controller to work
- Status updates require separate RBAC permissions
- Finalizers are crucial for proper cleanup

## AI Assistance Guidelines

### For Code Generation
1. Use Go best practices and controller-runtime patterns
2. Include proper error handling and logging
3. Add godoc comments for public functions
4. Consider test coverage for new code

### For Troubleshooting
1. Check generated manifests in `config/` directory
2. Look at RBAC permissions in comments
3. Consider Kubernetes API version compatibility
4. Check finalizer implementation for cleanup issues

### For Feature Implementation
1. Start with CRD changes if needed
2. Update controller logic incrementally
3. Add proper tests for new functionality
4. Update documentation and examples

## Example Conversation Starters

**For beginners**: "I'm new to Kubernetes operators. Can you explain how this MyNginx controller works and what it's supposed to do?"

**For implementation**: "I need to implement the missing ConfigMap feature. Can you help me update the CRD and controller logic?"

**For debugging**: "The controller isn't creating deployments properly. Can you help me debug the reconciliation logic?"

**For testing**: "How do I write proper tests for this controller? I need both unit and integration tests."

**For deployment**: "Help me deploy this controller to a Kubernetes cluster and test it with sample resources."
