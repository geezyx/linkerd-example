Example showing an issue I ran into where using the provided generated code from the linkerd/linkerd2 package seems to be running into validation errors when working with AuthorizationPolicy and MeshTLSAuthentication objects.

Running `go run main.go` with a kubernetes config pointed to a cluster with linkerd installed results in a validation error:

```
panic: AuthorizationPolicy.policy.linkerd.io "example" is invalid: spec.requiredAuthenticationRefs: Required value
```

Running `kubectl  apply -f authorizationpolicy.yaml` works without a validation error.
