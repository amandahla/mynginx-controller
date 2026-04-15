# Goal

Configure the controller to watch Deployments it owns and rely on Kubernetes garbage collection for cleanup instead of explicit deletion. This lesson demonstrates ownership-based watches and simplifies finalizer logic.

# Concepts

**Owns() Watch**: A watch that triggers reconciliation when resources owned by the custom resource change. This allows the controller to react when Deployments are modified or deleted externally.

**Owner References**: Metadata that marks one resource as owned by another. Kubernetes uses owner references for garbage collection and to establish object hierarchies.

**Garbage Collection**: Kubernetes automatically deletes dependent resources when their owner is deleted. Controllers can rely on this instead of implementing manual cleanup.

**Watch Predicates**: Filters that control when watches trigger reconciliation. The `Owns()` method automatically filters events to only reconcile when owned resources change.

**Simplified Finalizers**: When relying on garbage collection, finalizers may not need to delete resources explicitly. They can focus on external cleanup like API calls or resource de-registration.

# Result

This lesson modified the controller to:
- Add `.Owns(&appsv1.Deployment{})` to the controller setup, making it watch Deployments owned by MyNginx resources
- Remove the `reconcileDelete()` function since garbage collection handles Deployment deletion automatically
- Simplify the finalizer to only remove itself without explicitly deleting the Deployment
- Remove the manual `Get()` call for the Deployment in the main reconcile loop since watches provide notifications
- Clean up namespace resolution logic by using `myNginx.Namespace` directly

The controller now reacts to Deployment changes (like external deletion or modification) and recreates them as needed, without manual deletion logic.

# References

- [Controller Watches](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/builder#Builder.Owns)
- [Owner References](https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/)
- [Garbage Collection](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)
- [SetControllerReference](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/controller/controllerutil#SetControllerReference)
