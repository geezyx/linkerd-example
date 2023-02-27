Example showing an issue I ran into where using the provided generated code from the linkerd/linkerd2 package seems to be running into validation errors when working with AuthorizationPolicy and MeshTLSAuthentication objects.

Running `go run concrete/main.go` with a kubernetes config pointed to a cluster with linkerd installed results in a validation error:

```
panic: AuthorizationPolicy.policy.linkerd.io "example" is invalid: spec.requiredAuthenticationRefs: Required value
```

Using a dynamic client with unstructured objects seems to work around the issue.

Running `go run dynamic/main.go` with a kubernetes config pointed to a cluster with linkerd installed works without a validation error.

And finally, using a yaml file with `kubectl` directly also works as expected.

Running `kubectl  apply -f authorizationpolicy.yaml` with a kubernetes config pointed to a cluster with linkerd installed works without a validation error.
