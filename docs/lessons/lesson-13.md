# Goal

Add status conditions to report the operational state of the custom resource and implement volume mounts to connect the ConfigMap volume to the nginx container. This lesson covers Kubernetes status reporting patterns and container configuration.

# Concepts

**Status Conditions**: A standard Kubernetes pattern for reporting resource state. Conditions have a type (like "Ready"), status (True/False/Unknown), reason, and message. They provide structured observability.

**metav1.Condition**: A standard type for status conditions that includes Type, Status, Reason, Message, LastTransitionTime, and ObservedGeneration fields.

**VolumeMount**: Specification of how to mount a volume into a container. It defines the mount path and references a volume by name. Multiple containers can mount the same volume at different paths.

**Status Subresource**: A Kubernetes feature that separates the status field from the spec, allowing status updates without triggering spec validation. Enabled with `+kubebuilder:subresource:status`.

**Structured Status**: Organizing status information into well-defined fields instead of free-form messages makes it machine-readable and queryable.

# Result

This lesson implemented:
- A `Conditions` field in MyNginxStatus as a slice of `metav1.Condition`
- Condition types like "Ready" to indicate whether the Deployment is available
- Status updates in the reconciler using `meta.SetStatusCondition()` to report reconciliation outcomes
- VolumeMount configuration that mounts the ConfigMap volume at `/usr/share/nginx/html/` in the nginx container
- Conditional logic that only adds the volume mount when a ConfigMap is specified
- CRD annotations (`+listType=map`, `+listMapKey=type`) for proper condition handling

# References

- [Kubernetes API Conventions - Status](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)
- [Condition Type](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Condition)
- [Status Subresource](https://book.kubebuilder.io/reference/generating-crd.html#status)
- [Pod Container Configuration](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#Container)
