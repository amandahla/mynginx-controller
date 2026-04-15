# Goal

Replace the Update operation with Patch for modifying Deployment replicas and add a check to handle cases where the Deployment was externally deleted. This lesson covers advanced client operations and defensive programming.

# Concepts

**Patch vs Update**: Patch sends only the changed fields to the API server, while Update sends the entire object. Patch operations are more efficient and avoid conflicts when multiple controllers modify different fields.

**MergeFrom Patch**: A patch type that computes the difference between two objects. `client.MergeFrom(original)` creates a patch that captures changes made to the object after the original was copied.

**DeepCopy**: Creates an independent copy of an object including all nested fields. This is necessary to create a snapshot before modifications for patch comparison.

**Defensive Programming**: Checking error conditions (like `IsNotFound`) before performing operations that assume a resource exists prevents crashes and handles edge cases gracefully.

# Result

This lesson modified the update logic to:
- Use `client.Patch()` instead of `client.Update()` when changing Deployment replicas
- Create a patch using `client.MergeFrom(myDeployment.DeepCopy())` before modifying the Deployment
- Check if the Deployment was deleted (`IsNotFound`) in the finalizer before attempting to delete it
- Add proper error handling for the patch operation

# References

- [Kubernetes API Conventions - Patch](https://kubernetes.io/docs/reference/using-api/api-concepts/#patch)
- [controller-runtime Patch](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client.Patch)
- [Strategic Merge Patch](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/)
