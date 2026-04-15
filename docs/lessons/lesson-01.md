# Goal

Set up a Go project structure with basic tooling for building a Kubernetes controller. This lesson establishes the repository layout, dependency management, and build automation needed for controller development.

# Concepts

**Go module**: A collection of Go packages versioned together. The `go.mod` file defines the module path and Go version.

**Makefile**: A build automation file that defines targets (tasks) such as `build`, `test`, and `run` to standardize development workflows.

**.gitignore**: A file that tells Git which files and directories to ignore. This prevents binaries, build artifacts, and temporary files from being tracked in version control.

# Result

This lesson created the initial project structure with:
- A Go module initialized at `github.com/amandahla/mynginx-controller` using Go 1.23.4
- A Makefile with targets for testing, formatting, building, and running the controller
- A `.gitignore` file configured for Go projects to exclude binaries, test artifacts, and coverage files
- A README documenting the controller requirements: manage nginx deployments based on CRD specifications including replicas, ConfigMap for content, and Secret for TLS

# References

- [Go Modules Reference](https://go.dev/ref/mod)
- [GNU Make Manual](https://www.gnu.org/software/make/manual/make.html)
- [Kubernetes Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
