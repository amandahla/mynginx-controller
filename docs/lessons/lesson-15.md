# Goal

Improve code quality by extracting magic strings into constants, adding structured logging, and simplifying conditional logic. This lesson covers Go best practices for maintainable code.

# Concepts

**Constants**: Named values that do not change during execution. Constants make code self-documenting and prevent typos when the same value is used multiple times.

**Structured Logging**: Adding key-value pairs to log messages for filtering and querying. The `WithValues()` method enriches logs with context like resource name and namespace.

**Early Returns**: Returning from a function as soon as the result is known. This reduces nesting and makes the happy path easier to follow.

**Return Value Optimization**: Some functions like `AddFinalizer()` return a boolean indicating whether a change was made. Using this return value eliminates redundant checks.

# Result

This lesson refactored the code to:
- Define `MyNginxFinalizer` as a constant instead of repeating the string `"webapp.mynginx.amandahla.xyz/finalizer"` throughout the code
- Add structured logging with `.WithValues("name", req.Name, "namespace", req.Namespace)` to include resource context in all log messages
- Use the return value from `controllerutil.AddFinalizer()` to conditionally update only when the finalizer was actually added
- Simplify error handling by removing nested if statements and using early returns
- Combine multiple return statements into single lines where appropriate

# References

- [Effective Go - Constants](https://go.dev/doc/effective_go#constants)
- [Structured Logging in Go](https://github.com/go-logr/logr)
- [Early Return Pattern](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
