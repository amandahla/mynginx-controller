# Goal

Change from creating a Pod to creating a Deployment. Read the replicas field from the custom resource spec and use it to configure the Deployment. This lesson demonstrates managing higher-level workload resources.

# Concepts

**Deployment**: A Kubernetes resource that manages a ReplicaSet, which in turn manages Pods. Deployments provide declarative updates, rollback capabilities, and scaling.

**ReplicaSet**: A resource that maintains a stable set of replica Pods running at any given time. Deployments manage ReplicaSets automatically.

**Label Selector**: A mechanism to identify a set of Kubernetes objects. Deployments use label selectors to find the Pods they manage.

**Pod Template**: A specification for creating Pods. Deployments use the template to create Pods with identical specifications.

**ptr.To()**: A helper function that returns a pointer to a value. Kubernetes API fields that are optional or have zero values use pointers to distinguish between "not set" and "set to zero".

# Result

This lesson modified the reconciler to:
- Create a Deployment instead of a single Pod
- Read the `Replicas` value from the MyNginx spec and apply it to the Deployment
- Add labels to Pods (`app: <mynginx-name>`) so the Deployment can select and manage them
- Use `ptr.To(int32(myNginx.Spec.Replicas))` to convert the replica count to the required pointer type
- Set the controller reference so the Deployment is owned by the MyNginx resource

# References

- [Kubernetes Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)
- [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
- [API Conventions - Optional vs Required](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#optional-vs-required)
