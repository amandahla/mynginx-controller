# Goal

Create a Custom Resource Definition (CRD) and implement a basic reconciler that creates a Pod when the custom resource is created. This lesson introduces the core reconciliation loop pattern used by Kubernetes controllers.

# Concepts

**Custom Resource Definition (CRD)**: An extension to the Kubernetes API that defines a new resource type. CRDs allow you to create custom objects that behave like built-in Kubernetes resources.

**API Group**: A collection of related resources in Kubernetes. This controller uses `webapp.mynginx.amandahla.xyz` as the API group, which organizes the custom resources under a domain-specific namespace.

**Reconciler**: A function that runs whenever a resource changes. It compares the desired state (in the Spec) with the actual cluster state and takes actions to make them match.

**Controller Reference**: A relationship that marks one resource as owned by another. When the owner is deleted, Kubernetes automatically deletes the owned resources (garbage collection).

**RBAC markers**: Special comments (`+kubebuilder:rbac`) that generate RBAC permissions. Controllers need permissions to read and modify resources they manage.

**Scheme**: A registry that maps Go types to Kubernetes GroupVersionKind. The scheme is used to encode and decode Kubernetes objects.

# Result

This lesson created:
- A CRD for `MyNginx` resources in `api/v1/mynginx_types.go` with a `Replicas` field in the spec
- Auto-generated DeepCopy methods in `zz_generated.deepcopy.go` for object cloning (required by Kubernetes)
- A CRD manifest at `config/crd/bases/webapp.mynginx.amandahla.xyz_mynginxes.yaml`
- A reconciler in `internal/controller/mynginx_controller.go` that creates a single nginx Pod when a MyNginx resource is created
- RBAC rules granting the controller permissions to manage MyNginx resources
- Registration of the custom resource type in the controller manager's scheme
- A sample MyNginx custom resource in `config/samples/webapp_v1_mynginx.yaml`

# References

- [Kubernetes CRD Documentation](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)
- [controller-runtime Reconciler](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/reconcile)
- [Kubernetes Owner References](https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/)
- [Kubebuilder API Design](https://book.kubebuilder.io/cronjob-tutorial/api-design)
