# Goal

Improve the finalizer implementation and write comprehensive unit tests for the controller. This lesson introduces test-driven development practices for Kubernetes controllers.

# Concepts

**Unit Testing**: Testing individual components in isolation. Controller tests verify that reconciliation logic produces the expected results without requiring a real cluster.

**envtest**: A testing framework that runs a real API server and etcd for integration testing. It provides a realistic environment without the overhead of a full cluster.

**Test Suite**: A collection of related tests organized using a testing framework. Controller-runtime uses Ginkgo and Gomega for BDD-style tests.

**Assertions**: Statements that verify expected outcomes. Gomega provides expressive matchers like `Should()`, `Equal()`, and `BeNil()` for readable assertions.

**Test Lifecycle**: Setup and teardown operations that run before and after tests. Creating and deleting resources in tests requires cleanup to avoid state leakage.

# Result

This lesson:
- Simplified the finalizer logic to be more robust and easier to understand
- Created comprehensive unit tests in `mynginx_controller_test.go` that verify:
  - Custom resource creation triggers Deployment creation
  - Updating replicas in the custom resource updates the Deployment
  - Deleting the custom resource runs the finalizer and cleans up the Deployment
- Used envtest to run tests against a real API server
- Added test assertions using Gomega matchers for readable test expectations

# References

- [Kubebuilder Testing](https://book.kubebuilder.io/cronjob-tutorial/writing-tests.html)
- [envtest Documentation](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/envtest)
- [Ginkgo Testing Framework](https://onsi.github.io/ginkgo/)
- [Gomega Matcher Library](https://onsi.github.io/gomega/)
