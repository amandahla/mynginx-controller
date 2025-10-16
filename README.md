# mynginx-controller

## Requirements

- The controller should read a CRD with definitions of:
 - Number of replicas (default: 1)
 - Name of ConfigMap to use as content for the file "/usr/share/nginx/html/index.html" (optional)
 - Name of Secret to set TLS (optional)

- Then will create a deployment with NGINX image.

- If ConfigMap or Secret don't exist, update status field with failure message.
- If CRD changes, reconcile.
- If ConfigMap changes, ignore.
- If Secret changes, reconcile.

## Changelog

### 2025-09-19
Refactor finalizer and ensureDeployment.

### 2025-09-01
Controller creates, updates and deletes deployment based on CRD mynginx.
  
