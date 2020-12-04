# goelector

Lib that extends https://pkg.go.dev/k8s.io/client-go/tools/leaderelection

# Example

Example folder provides full build/deploy/test makefile targets

```
Usage:
  make <target>

Targets:
  compile Go compiles elector for linux os
  docker Docker build and push to your registry, needs REGISTRY env var to be set
  check Check whichs pod is elected leader
  deploy Deploy serviceaccount and deployment to k8s, requires REGISTRY env var to be set
  deploy-sa Deploys service account with access to coordination api, only needed once
  help Show help
```

## Docker

`$REGISTRY=gcr.io/<your_project> make docker`
Builds dockerfile, tags and pushes to registry

`$REGISTRY=gcr.io/<your_project> make deploy`
Deploys to active kubernetes cluster/namespace
Using extended service account created by `$make deploy-sa`

`$REGISTRY=gcr.io/<your_project> make check`
Checks of elector is selected, example output

```
Running services:
elector-test-68b9db5f95-64cvv
elector-test-68b9db5f95-h8db8
elector-test-68b9db5f95-wvtd9

Elected leader:
elector-test-68b9db5f95-h8db8
```

## Usage

### Start

Blocking, using only default functions for callbacks when leader is selected

### StartWithCallbacks

Blocking, provides user defined callback functions for Start/Stop/New Leader
