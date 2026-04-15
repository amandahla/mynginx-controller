# Goal

Implement test isolation by creating random namespaces for each test and cleaning them up afterward. This lesson covers test best practices that prevent tests from interfering with each other.

# Concepts

**Test Isolation**: Each test should run independently without affecting other tests. Shared state between tests causes flaky failures and makes debugging difficult.

**Random Namespace Generation**: Creating a unique namespace per test ensures resources from different tests do not collide. Random names prevent conflicts even when tests run in parallel.

**Test Cleanup**: Deleting resources after tests complete prevents resource accumulation and ensures a clean state for subsequent tests. The `AfterEach` hook runs cleanup code after each test.

**BeforeEach and AfterEach**: Lifecycle hooks that run code before and after each test. These hooks set up test fixtures and clean up resources automatically.

# Result

This lesson modified the test suite to:
- Generate a random namespace name in `BeforeEach` using `fmt.Sprintf("test-ns-%d", time.Now().UnixNano())`
- Create the namespace before each test runs
- Delete the namespace in `AfterEach` to clean up all resources created during the test
- Remove manual deletion of individual resources since namespace deletion cascades to all contained resources
- Simplified test code by centralizing setup and teardown logic

# References

- [Kubernetes Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
- [Ginkgo BeforeEach/AfterEach](https://onsi.github.io/ginkgo/#extracting-common-setup-beforeeach)
- [Test Isolation Best Practices](https://testing.googleblog.com/2015/04/just-say-no-to-more-end-to-end-tests.html)
