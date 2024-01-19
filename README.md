# Dagger Grype module

Known to work with Dagger v0.9.5.

Check containers for vulnerabilities using [Grype](https://github.com/anchore/grype) from your Dagger pipelines.

## Containers

Check the given container image for vulnerabilities:

### From Dagger CLI

```
dagger call -m github.com/lukemarsden/dagger-grype test-container --image "alpine:latest"
```

### From Dagger Code

Call the `TestContainer` method on this module with your container object (e.g. `*Container` in golang) as the only argument.
This allows you to test a container without first pushing it to a container registry.

Internally, this module exports the container image as a tarball and passes that tarball to grype.
