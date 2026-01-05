# Development Tasks & TODO

## High Priority (Missing Requirements)

### 1. ConfigMap Integration
- [ ] Add `ConfigMapName` field to MyNginxSpec
- [ ] Update controller to mount ConfigMap as index.html
- [ ] Add validation for ConfigMap existence
- [ ] Update status on ConfigMap not found
- [ ] Add tests for ConfigMap scenarios

### 2. Secret Integration (TLS)
- [ ] Add `SecretName` field to MyNginxSpec for TLS
- [ ] Update controller to mount Secret as TLS cert
- [ ] Add HTTPS port configuration
- [ ] Update status on Secret not found
- [ ] Watch Secret changes for reconciliation
- [ ] Add tests for TLS scenarios

### 3. Status Reporting
- [ ] Define MyNginxStatus fields (Ready, Error, Message)
- [ ] Update status on reconciliation success/failure
- [ ] Add status conditions pattern
- [ ] Report ConfigMap/Secret validation errors
- [ ] Add status-focused tests

## Medium Priority (Improvements)

### 4. Enhanced Reconciliation
- [ ] Watch ConfigMap changes (currently ignored per requirements)
- [ ] Improve Secret change detection
- [ ] Add proper deployment update strategies
- [ ] Implement progressive rollout support

### 5. Error Handling & Resilience
- [ ] Add retry logic with exponential backoff
- [ ] Better error categorization (temporary vs permanent)
- [ ] Add metrics for monitoring
- [ ] Improve logging with more context

### 6. Testing & Validation
- [ ] Expand unit test coverage
- [ ] Add comprehensive e2e tests
- [ ] Add chaos testing scenarios
- [ ] Validate CRD schema changes

## Low Priority (Nice-to-Have)

### 7. Advanced Features
- [ ] Support for multiple ConfigMaps
- [ ] Custom nginx configuration support
- [ ] Resource limits and requests configuration
- [ ] Pod disruption budget support
- [ ] HorizontalPodAutoscaler integration

### 8. Developer Experience
- [ ] Add more detailed documentation
- [ ] Create example manifests
- [ ] Add development environment setup scripts
- [ ] Improve debugging tools

### 9. Production Readiness
- [ ] Add health checks and readiness probes
- [ ] Implement graceful shutdown
- [ ] Add container security context
- [ ] Resource optimization

## Technical Debt

### Code Quality
- [ ] Refactor reconcileDeployment function (getting large)
- [ ] Extract deployment creation to separate function
- [ ] Add more godoc comments
- [ ] Improve error messages

### Build & CI/CD
- [ ] Add automated testing in CI
- [ ] Container image build automation
- [ ] Release automation
- [ ] Security scanning

### Documentation
- [ ] API documentation generation
- [ ] Operator lifecycle documentation
- [ ] Troubleshooting guide
- [ ] Migration guide for updates

## Current Implementation Gaps

Based on the requirements in README.md, these are missing:

1. **ConfigMap for HTML content** - Not implemented
2. **Secret for TLS** - Not implemented  
3. **Status field updates** - Basic structure exists but no logic
4. **ConfigMap change handling** - Should be ignored per requirements
5. **Secret change reconciliation** - Not implemented

## Quick Wins (Easy to implement)

- [ ] Add validation for replica count (> 0)
- [ ] Add defaulting for replica count (default: 1)
- [ ] Improve logging messages
- [ ] Add owner references for created resources
- [ ] Add labels for better resource management

## Architecture Improvements

- [ ] Consider using Helm charts for deployment
- [ ] Add admission webhooks for validation
- [ ] Implement conversion webhooks for API versioning
- [ ] Add custom metrics for monitoring
