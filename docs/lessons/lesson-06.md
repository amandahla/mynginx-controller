# Goal

Implement a finalizer to ensure the Deployment is properly deleted when the MyNginx custom resource is deleted. This lesson covers cleanup logic and preventing resource leaks.

# Concepts

**Finalizer**: A string value in the resource metadata that prevents deletion until cleanup tasks complete. Kubernetes blocks deletion of resources with finalizers until controllers remove them.

**DeletionTimestamp**: A field set by Kubernetes when a delete request is received. Resources with finalizers enter a "being deleted" state instead of being immediately removed.

**Garbage Collection**: Kubernetes automatically deletes owned resources when the owner is deleted, but finalizers allow custom cleanup logic to run before deletion completes.

**controllerutil**: A package from controller-runtime that provides helper functions for common controller operations like adding and removing finalizers.

# Result

This lesson implemented finalizer logic that:
- Defines a finalizer named `webapp.mynginx.amandahla.xyz/finalizer`
- Adds the finalizer to new MyNginx resources during reconciliation
- Checks if the resource is being deleted by examining the `DeletionTimestamp`
- Deletes the Deployment explicitly when the MyNginx resource is being deleted
- Removes the finalizer after cleanup completes, allowing Kubernetes to finish deletion
- Stops reconciliation early if the resource is being deleted

# References

- [Kubernetes Finalizers](https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers/)
- [Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html)
- [controllerutil Package](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/controller/controllerutil)
