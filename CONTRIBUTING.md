# Contributing Guidelines

## Getting Started

### Prerequisites
- Go 1.23.0+
- Docker
- Kubernetes cluster (or kind/minikube for local development)
- kubectl configured

### Development Setup
1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Install development tools: `make install-tools`
4. Run tests: `make test`
5. Run locally: `make run`

## Code Standards

### Go Code Guidelines
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add godoc comments for public functions and types
- Handle errors explicitly
- Use structured logging with context

### Kubernetes Controller Guidelines
- Follow controller-runtime patterns
- Use proper RBAC annotations
- Implement finalizers for cleanup
- Use structured logging with request context
- Handle not-found errors appropriately

### Testing
- Write unit tests for business logic
- Add integration tests for controller behavior
- Use Ginkgo/Gomega for consistent test structure
- Test both success and error scenarios

## Development Workflow

### Making Changes
1. Create a feature branch from main
2. Make your changes
3. Run `make` to generate code and run tests
4. Test locally with `make run`
5. Create pull request

### CRD Changes
1. Modify types in `api/v1/mynginx_types.go`
2. Run `make` to regenerate code
3. Update controller logic if needed
4. Add/update tests
5. Update documentation

### Adding New Features
1. Design the feature (consider CRD changes, controller logic, tests)
2. Update the CRD schema if needed
3. Implement controller changes
4. Add proper RBAC permissions
5. Write tests
6. Update documentation

## Project Structure

```
mynginx-controller/
├── api/v1/                 # API definitions
├── cmd/                    # Application entrypoints
├── config/                 # Deployment configurations
├── internal/controller/    # Controller implementation
├── test/                   # Test files
├── hack/                   # Development scripts
└── Makefile               # Build automation
```

## Common Tasks

### Running Tests
```bash
make test                  # Run unit tests
make test-e2e             # Run end-to-end tests
make envtest              # Set up test environment
```

### Building and Deploying
```bash
make build                 # Build binary
make docker-build         # Build container image
make deploy               # Deploy to cluster
make undeploy             # Remove from cluster
```

### Debugging
```bash
make run ARGS="--log-level=debug"  # Run with debug logging
kubectl logs -f deployment/mynginx-controller-manager -n mynginx-controller-system
```

## Pull Request Guidelines

### Before Submitting
- [ ] Tests pass (`make test`)
- [ ] Code is formatted (`gofmt`)
- [ ] Generated code is up to date (`make`)
- [ ] Documentation is updated
- [ ] RBAC permissions are correct

### PR Description Should Include
- What changes were made
- Why the changes were needed
- How to test the changes
- Any breaking changes
- Related issues/tickets

## Release Process

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create release branch
4. Test thoroughly
5. Tag release
6. Build and push container image
7. Create GitHub release

## Getting Help

- Check existing documentation in `/docs`
- Look at similar implementations in controller-runtime examples
- Ask questions in issues or discussions
- Check Kubernetes Slack channels for controller-runtime help
