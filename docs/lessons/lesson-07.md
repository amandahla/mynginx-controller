# Goal

Refactor the Reconcile function by extracting logic into separate helper functions. This lesson demonstrates code organization practices that make controllers easier to read, test, and maintain.

# Concepts

**Function Extraction**: Breaking a large function into smaller, focused functions. Each function should have a single responsibility and a clear name that describes what it does.

**Separation of Concerns**: Organizing code so different aspects of the logic (finalizer handling, deployment creation, namespace resolution) are isolated in dedicated functions.

**Code Readability**: Well-structured code with descriptive function names acts as documentation, making it easier for developers to understand the control flow.

# Result

This lesson refactored the reconciler by extracting three helper functions:
- `handleFinalizer()`: Manages finalizer logic for both adding and cleanup during deletion
- `ensureDeployment()`: Handles deployment creation and updates
- Helper logic for namespace resolution moved into focused sections

The main `Reconcile()` function now reads as a high-level workflow: fetch the resource, handle finalizers, and ensure the deployment exists. The implementation details are delegated to helper functions.

# References

- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Clean Code Principles](https://github.com/ryanmcdermott/clean-code-javascript#functions)
- [Refactoring Guru](https://refactoring.guru/extract-method)
