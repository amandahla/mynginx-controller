# Goal

Add Kubernetes event recording to the controller to communicate actions and errors to users. This lesson introduces the Events API and demonstrates how controllers surface operational information through Kubernetes events visible in `kubectl describe`.

# Concepts

**Event**: A Kubernetes resource that records controller actions, state changes, warnings, and errors. Events appear in `kubectl describe` output and help users debug resource issues.

**EventRecorder**: An interface that creates Event resources associated with a Kubernetes object. Controllers use EventRecorder to publish events about the resources they manage.

**Event Types**: Events are categorized as Normal (informational) or Warning (problems). The type helps users filter and prioritize event notifications.

**Event Reason**: A short, machine-readable string (like `ConfigMapNotFound`, `DeploymentCreated`) that categorizes the event. Reasons should be CamelCase constants.

**Event Action**: A label describing what operation triggered the event (like `CreateDeployment`, `ChangeDeployment`). Actions provide additional context about the controller's behavior.

**RBAC for Events**: Controllers need explicit permissions to create Event resources. The RBAC marker grants access to the events API group.

# Result

This lesson added event recording throughout the controller:
- Imported `k8s.io/client-go/tools/events` package for event recording capabilities
- Added `Recorder events.EventRecorder` field to `MyNginxReconciler` struct
- Injected the recorder in `cmd/main.go` using `mgr.GetEventRecorder("mynginx-controller")`
- Added RBAC marker `+kubebuilder:rbac:groups=events.k8s.io,resources=events,verbs=create;patch` for event permissions
- Published Warning event when ConfigMap is not found during Deployment creation
- Published Normal event when Deployment is successfully created
- Published Normal event when Deployment replicas are changed
- Bumped Go version from 1.23.0 to 1.26.1
- Updated Kubernetes dependencies from v0.32.1 to v0.35.0
- Updated controller-runtime from v0.20.2 to v0.23.3

# References

- [Kubernetes Events](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/event-v1/)
- [Kubebuilder Raising Events](https://book.kubebuilder.io/reference/raising-events.html)
- [EventRecorder Interface](https://pkg.go.dev/k8s.io/client-go/tools/record#EventRecorder)
- [Event API Types](https://pkg.go.dev/k8s.io/api/core/v1#Event)
- [Go Modules - Updating Dependencies](https://go.dev/doc/modules/managing-dependencies#getting_version)
