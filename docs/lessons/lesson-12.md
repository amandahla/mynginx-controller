# Goal

Add ConfigMap support to the custom resource spec, allowing users to provide custom nginx content. This lesson demonstrates how to extend the CRD with optional configuration and mount ConfigMaps as volumes in Pods.

# Concepts

**ConfigMap**: A Kubernetes resource that stores configuration data as key-value pairs. ConfigMaps decouple configuration from container images, making applications portable.

**Volume**: An abstraction that represents a directory accessible to containers in a Pod. Volumes can be backed by various sources including ConfigMaps, Secrets, and persistent storage.

**Volume Source**: The backing store for a volume. `ConfigMapVolumeSource` creates a volume from a ConfigMap, making its data available as files in the container filesystem.

**Optional Fields**: Fields in the spec that users may omit. The `omitempty` JSON tag prevents the field from appearing in YAML when not set.

# Result

This lesson added:
- An `IndexConfigMapName` field to the MyNginxSpec for specifying a ConfigMap containing custom HTML content
- Logic to create a Volume from the ConfigMap when the field is set
- A conditional check that only adds the volume if `IndexConfigMapName` is not empty
- Updated CRD manifest with the new field and documentation
- Code that constructs the PodSpec separately to conditionally include volumes

# References

- [Kubernetes ConfigMaps](https://kubernetes.io/docs/concepts/configuration/configmap/)
- [Volumes](https://kubernetes.io/docs/concepts/storage/volumes/)
- [Configure a Pod to Use a ConfigMap](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/)
