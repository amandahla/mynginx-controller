# Goal

Change the Replicas field type from `int` to `int32` to match Kubernetes API conventions. This lesson covers API compatibility and type alignment with standard Kubernetes resources.

# Concepts

**int32 vs int**: Kubernetes uses `int32` for numeric fields because the API is designed to work across different architectures. Using `int32` ensures consistent behavior on 32-bit and 64-bit systems.

**API Compatibility**: Custom resource types should follow the same conventions as built-in Kubernetes types. Deployment replicas use `*int32`, so the custom resource should also use `int32`.

**Type Conversion**: When types change, you must convert values explicitly. This prevents implicit conversions that might hide bugs.

**CRD Schema Updates**: Changing field types in Go structs requires regenerating the CRD manifests using `make manifests` so the OpenAPI schema matches.

# Result

This lesson changed:
- The `Replicas` field in `MyNginxSpec` from `int` to `int32`
- Removed `ptr.To()` calls and type conversions since both the custom resource and Deployment now use the same type
- Updated the CRD manifest to reflect the schema change
- Simplified code by eliminating unnecessary type conversions

# References

- [Kubernetes API Conventions - Types](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds)
- [Go Integer Types](https://go.dev/ref/spec#Numeric_types)
- [Deployment Spec](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/#DeploymentSpec)
