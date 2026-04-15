# Goal

Extract Deployment creation logic into a dedicated function to improve code organization and reduce duplication. This lesson demonstrates the Single Responsibility Principle in controller design.

# Concepts

**Single Responsibility Principle**: Each function should do one thing well. Separating deployment creation from deployment management makes both operations easier to understand and test.

**Function Signatures**: Well-designed functions have clear inputs and outputs. `createDeployment(ctx, myNginx)` clearly indicates it creates a Deployment from a MyNginx resource.

**Code Reusability**: Extracting creation logic allows it to be reused or tested independently. The function can be called from multiple places or mocked in tests.

**Separation of Concerns**: The `ensureDeployment` function now handles the decision logic (create vs update), while `createDeployment` handles the creation details.

# Result

This lesson created the `createDeployment()` function that:
- Takes a MyNginx resource as input and returns an error
- Builds the complete Deployment specification with labels, replica count, and pod template
- Sets the controller reference to establish ownership
- Creates the Deployment in the cluster

The `ensureDeployment()` function now delegates to `createDeployment()` when the Deployment does not exist, making the code cleaner and the responsibilities clearer.

# References

- [Single Responsibility Principle](https://en.wikipedia.org/wiki/Single-responsibility_principle)
- [Clean Code - Functions](https://github.com/ryanmcdermott/clean-code-javascript#functions)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
