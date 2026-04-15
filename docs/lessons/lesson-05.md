# Goal

Add update logic to detect when the replica count in the custom resource differs from the Deployment and update the Deployment to match. This lesson introduces the reconciliation pattern of comparing desired state with actual state.

# Concepts

**Reconciliation Loop**: The core pattern where a controller continuously compares the desired state (in the custom resource) with the actual state (in the cluster) and takes actions to align them.

**Update vs Create**: Controllers must handle both creating new resources and updating existing ones. The `Update()` call modifies an existing resource in the cluster.

**Idempotency**: Controllers should produce the same result regardless of how many times reconciliation runs. Checking before updating ensures idempotent behavior.

# Result

This lesson added logic to:
- Compare the current Deployment replica count with the desired count from the MyNginx spec
- Update the Deployment when the replica counts differ
- Log when an update occurs for observability

The reconciler now handles two scenarios: creating a new Deployment if it does not exist, and updating an existing Deployment if the replica count changes.

# References

- [Controller Pattern](https://kubernetes.io/docs/concepts/architecture/controller/)
- [Level Triggering vs Edge Triggering](https://hackernoon.com/level-triggering-and-reconciliation-in-kubernetes-1f17fe30333d)
- [Client Update Method](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client.Update)
