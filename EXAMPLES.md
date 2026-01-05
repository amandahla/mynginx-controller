# MyNginx Controller - Examples and Usage

## Quick Start

### 1. Deploy the Controller
```bash
# Build and deploy
make deploy

# Verify deployment
kubectl get deployment -n mynginx-controller-system
```

### 2. Create a Basic MyNginx Resource
```yaml
# examples/basic-mynginx.yaml
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: my-nginx-sample
  namespace: default
spec:
  replicas: 3
```

```bash
kubectl apply -f examples/basic-mynginx.yaml
kubectl get mynginx
kubectl get deployment
```

## Example Scenarios

### Basic Nginx Deployment
```yaml
# examples/simple.yaml
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: simple-nginx
spec:
  replicas: 2
---
# Check the created deployment
# kubectl get deployment simple-nginx
```

### With Custom HTML Content (Future Implementation)
```yaml
# examples/with-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-html-content
data:
  index.html: |
    <!DOCTYPE html>
    <html>
    <head><title>My Custom Page</title></head>
    <body>
      <h1>Hello from MyNginx Controller!</h1>
      <p>This content comes from a ConfigMap.</p>
    </body>
    </html>
---
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: custom-content-nginx
spec:
  replicas: 1
  configMapName: my-html-content  # Not yet implemented
```

### With TLS Certificate (Future Implementation)
```yaml
# examples/with-tls.yaml
apiVersion: v1
kind: Secret
metadata:
  name: nginx-tls
type: kubernetes.io/tls
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0t...  # base64 encoded cert
  tls.key: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0t...  # base64 encoded key
---
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: secure-nginx
spec:
  replicas: 2
  secretName: nginx-tls  # Not yet implemented
```

### Complete Example (Future Implementation)
```yaml
# examples/complete.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-content
data:
  index.html: |
    <!DOCTYPE html>
    <html>
    <head><title>Secure App</title></head>
    <body>
      <h1>Secure Application</h1>
      <p>Served over HTTPS with custom content!</p>
    </body>
    </html>
---
apiVersion: v1
kind: Secret
metadata:
  name: app-tls
type: kubernetes.io/tls
stringData:
  tls.crt: |
    -----BEGIN CERTIFICATE-----
    # Your certificate here
    -----END CERTIFICATE-----
  tls.key: |
    -----BEGIN PRIVATE KEY-----
    # Your private key here
    -----END PRIVATE KEY-----
---
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: complete-app
spec:
  replicas: 3
  configMapName: app-content  # Not yet implemented
  secretName: app-tls         # Not yet implemented
```

## Development Examples

### Local Development
```bash
# Run controller locally (requires kubeconfig)
make run

# In another terminal, apply test resources
kubectl apply -f examples/basic-mynginx.yaml

# Watch logs for reconciliation
```

### Testing Changes
```bash
# After making code changes
make test

# Test deployment creation
kubectl apply -f examples/basic-mynginx.yaml
kubectl describe deployment simple-nginx

# Test deletion and cleanup
kubectl delete mynginx simple-nginx
# Should clean up deployment due to finalizer
```

### Debugging
```bash
# Check controller logs
kubectl logs -f deployment/mynginx-controller-manager -n mynginx-controller-system

# Check resource status
kubectl describe mynginx my-nginx-sample

# Check created deployment
kubectl describe deployment my-nginx-sample
```

## Common Use Cases

### 1. Simple Web Application Hosting
Deploy static websites with configurable scaling:
```yaml
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: website-prod
spec:
  replicas: 5
```

### 2. Development Environment
Quick nginx instances for testing:
```yaml
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: dev-test
spec:
  replicas: 1
```

### 3. Load Balancer Frontend (Future)
With TLS termination and custom content:
```yaml
# When ConfigMap/Secret features are implemented
apiVersion: webapp.mynginx.amandahla.xyz/v1
kind: MyNginx
metadata:
  name: lb-frontend
spec:
  replicas: 3
  configMapName: frontend-config
  secretName: frontend-tls
```

## Monitoring and Operations

### Check Resource Status
```bash
# List all MyNginx resources
kubectl get mynginx

# Detailed view
kubectl describe mynginx <name>

# Check associated deployments
kubectl get deployment -l app=mynginx
```

### Scaling Operations
```bash
# Scale up/down by editing the resource
kubectl patch mynginx my-nginx-sample -p '{"spec":{"replicas":5}}'

# Or edit directly
kubectl edit mynginx my-nginx-sample
```

### Cleanup
```bash
# Delete specific resource
kubectl delete mynginx my-nginx-sample

# Delete all MyNginx resources
kubectl delete mynginx --all

# Uninstall controller
make undeploy
```

## Troubleshooting

### Controller Not Creating Deployments
1. Check controller logs
2. Verify RBAC permissions
3. Check CRD is properly installed

### Deployment Not Starting
1. Check deployment events: `kubectl describe deployment <name>`
2. Check pod logs: `kubectl logs <pod-name>`
3. Verify image availability

### Controller Crashes
1. Check controller manager logs
2. Verify Kubernetes API compatibility
3. Check for resource conflicts
